package fraction

import (
	"fmt"
    "strconv"
	"strings"
)

type Fraction struct {
    numerator, denominator int64
}

func (frac Fraction) Numerator() int64 {
    return frac.numerator
}

func (frac Fraction) Denominator() int64 {
    return frac.denominator
}

func (frac Fraction) String() string {
    return fmt.Sprintf("%d/%d", frac.numerator, frac.denominator)
}

func (frac Fraction) MultByFraction(other Fraction) Fraction {
    return Fraction {
        numerator: frac.numerator * other.numerator,
        denominator: frac.denominator * other.denominator,
    }
}

func (frac Fraction) MultByInt(other int) Fraction {
    return Fraction {
        numerator: frac.numerator * int64(other),
        denominator: frac.denominator,
    }
}

func (frac Fraction) IsInt() bool {
    return frac.denominator % frac.numerator == 0
}

func ReadFraction(text string) (frac Fraction, ok bool) {
    text = strings.TrimSpace(text)
    parts := strings.Split(text, "/")
    if len(parts) != 2 {
        return Fraction{}, false
    }
    var err error
    frac.numerator, err = strconv.ParseInt(parts[0], 10, 64)
    if err != nil {
        return Fraction{}, false
    }
    frac.denominator, err = strconv.ParseInt(parts[1], 10, 64)
    if err != nil {
        return Fraction{}, false
    }
    return frac, true
}
