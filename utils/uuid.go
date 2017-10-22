package utils

import (
	randc "crypto/rand"
	"errors"
	"fmt"
	"math"
	randm "math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

)

// DefaultABC
const DefaultABC = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Abc struct {
	alphabet []rune
}

type Shortid struct {
	abc    Abc
	worker uint
	epoch  time.Time
	ms     uint
	count  uint
	mx     sync.Mutex
}

var shortid *Shortid

func init() {
	shortid = MustNew(0, DefaultABC, 1)
}

func GetDefault() *Shortid {
	return (*Shortid)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&shortid))))
}

func SetDefault(sid *Shortid) {
	target := (*unsafe.Pointer)(unsafe.Pointer(&shortid))
	source := unsafe.Pointer(sid)
	atomic.SwapPointer(target, source)
}

func Generate() (string, error) {
	return shortid.Generate()
}

func MustGenerate() string {
	id, err := Generate()
	if err == nil {
		return id
	}
	panic(err)
}

func New(worker uint8, alphabet string, seed uint64) (*Shortid, error) {
	if worker > 31 {
		return nil, errors.New("expected worker in the range [0,31]")
	}
	abc, err := NewAbc(alphabet, seed)
	if err == nil {
		sid := &Shortid{
			abc:    abc,
			worker: uint(worker),
			epoch:  time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			ms:     0,
			count:  0,
		}
		return sid, nil
	}
	return nil, err
}

func MustNew(worker uint8, alphabet string, seed uint64) *Shortid {
	sid, err := New(worker, alphabet, seed)
	if err == nil {
		return sid
	}
	panic(err)
}

func (sid *Shortid) Generate() (string, error) {
	return sid.GenerateInternal(nil, sid.epoch)
}

func (sid *Shortid) MustGenerate() string {
	id, err := sid.Generate()
	if err == nil {
		return id
	}
	panic(err)
}

func (sid *Shortid) GenerateInternal(tm *time.Time, epoch time.Time) (string, error) {
	ms, count := sid.getMsAndCounter(tm, epoch)
	idrunes := make([]rune, 9)
	if tmp, err := sid.abc.Encode(ms, 8, 5); err == nil {
		copy(idrunes, tmp)
	} else {
		return "", err
	}
	if tmp, err := sid.abc.Encode(sid.worker, 1, 5); err == nil {
		idrunes[8] = tmp[0]
	} else {
		return "", err
	}
	if count > 0 {
		if countrunes, err := sid.abc.Encode(count, 0, 6); err == nil {
			idrunes = append(idrunes, countrunes...)
		} else {
			return "", err
		}
	}
	return string(idrunes), nil
}

func (sid *Shortid) getMsAndCounter(tm *time.Time, epoch time.Time) (uint, uint) {
	sid.mx.Lock()
	defer sid.mx.Unlock()
	var ms uint
	if tm != nil {
		ms = uint(tm.Sub(epoch).Nanoseconds() / 1000000)
	} else {
		ms = uint(time.Now().Sub(epoch).Nanoseconds() / 1000000)
	}
	if ms == sid.ms {
		sid.count++
	} else {
		sid.count = 0
		sid.ms = ms
	}
	return sid.ms, sid.count
}

func (sid *Shortid) String() string {
	return fmt.Sprintf("Shortid(worker=%v, epoch=%v, abc=%v)", sid.worker, sid.epoch, sid.abc)
}

func (sid *Shortid) Abc() Abc {
	return sid.abc
}

func (sid *Shortid) Epoch() time.Time {
	return sid.epoch
}

func (sid *Shortid) Worker() uint {
	return sid.worker
}

func NewAbc(alphabet string, seed uint64) (Abc, error) {
	runes := []rune(alphabet)
	if len(runes) != len(DefaultABC) {
		return Abc{}, fmt.Errorf("alphabet must contain %v unique characters", len(DefaultABC))
	}
	if nonUnique(runes) {
		return Abc{}, errors.New("alphabet must contain unique characters only")
	}
	abc := Abc{alphabet: nil}
	abc.shuffle(alphabet, seed)
	return abc, nil
}

func MustNewAbc(alphabet string, seed uint64) Abc {
	res, err := NewAbc(alphabet, seed)
	if err == nil {
		return res
	}
	panic(err)
}

func nonUnique(runes []rune) bool {
	found := make(map[rune]struct{})
	for _, r := range runes {
		if _, seen := found[r]; !seen {
			found[r] = struct{}{}
		}
	}
	return len(found) < len(runes)
}

func (abc *Abc) shuffle(alphabet string, seed uint64) {
	source := []rune(alphabet)
	for len(source) > 1 {
		seed = (seed*9301 + 49297) % 233280
		i := int(seed * uint64(len(source)) / 233280)

		abc.alphabet = append(abc.alphabet, source[i])
		source = append(source[:i], source[i+1:]...)
	}
	abc.alphabet = append(abc.alphabet, source[0])
}

func (abc *Abc) Encode(val, nsymbols, digits uint) ([]rune, error) {
	if digits < 4 || 6 < digits {
		return nil, fmt.Errorf("allowed digits range [4,6], found %v", digits)
	}

	var computedSize uint = 1
	if val >= 1 {
		computedSize = uint(math.Log2(float64(val)))/digits + 1
	}
	if nsymbols == 0 {
		nsymbols = computedSize
	} else if nsymbols < computedSize {
		return nil, fmt.Errorf("cannot accommodate data, need %v digits, got %v", computedSize, nsymbols)
	}

	mask := 1<<digits - 1

	random := make([]int, int(nsymbols))

	if digits < 6 {
		copy(random, maskedRandomInts(len(random), 0x3f-mask))
	}

	res := make([]rune, int(nsymbols))
	for i := range res {
		shift := digits * uint(i)
		index := (int(val>>shift) & mask) | random[i]
		res[i] = abc.alphabet[index]
	}
	return res, nil
}

func (abc *Abc) MustEncode(val, size, digits uint) []rune {
	res, err := abc.Encode(val, size, digits)
	if err == nil {
		return res
	}
	panic(err)
}

func maskedRandomInts(size, mask int) []int {
	ints := make([]int, size)
	bytes := make([]byte, size)
	if _, err := randc.Read(bytes); err == nil {
		for i, b := range bytes {
			ints[i] = int(b) & mask
		}
	} else {
		for i := range ints {
			ints[i] = randm.Intn(0xff) & mask
		}
	}
	return ints
}

func (abc Abc) String() string {
	return fmt.Sprintf("Abc{alphabet='%v')", abc.Alphabet())
}

func (abc Abc) Alphabet() string {
	return string(abc.alphabet)
}
