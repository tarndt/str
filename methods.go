package str

import (
	"strings"
	"unicode"
)

func (s String) Clone() String {
	return String(strings.Clone(string(s)))
}

func (s String) Compare(other String) int {
	return strings.Compare(string(s), string(other))
}

func (s String) Contains(subStr String) bool {
	return strings.Contains(string(s), string(subStr))
}

func (s String) ContainsAny(chars String) bool {
	return strings.ContainsAny(string(s), string(chars))
}

func (s String) ContainsRune(r rune) bool {
	return strings.ContainsRune(string(s), r)
}

func (s String) Count(subStr String) int {
	return strings.Count(string(s), string(subStr))
}

func (s String) Cut(sep String) (before, after String, found bool) {
	prefix, suffix, found := strings.Cut(string(s), string(sep))
	return String(prefix), String(suffix), found
}

func (s String) EqualFold(t String) bool {
	return strings.EqualFold(string(s), string(t))
}

func (s String) Fields() Strings {
	return Slice(strings.Fields(string(s))...)
}

func (s String) FieldsFunc(f func(rune) bool) Strings {
	return Slice(strings.FieldsFunc(string(s), f)...)
}

func (s String) HasPrefix(prefix String) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

func (s String) HasSuffix(suffix String) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

func (s String) Index(subStr String) int {
	return strings.Index(string(s), string(subStr))
}

func (s String) IndexAny(chars String) int {
	return strings.IndexAny(string(s), string(chars))
}

func (s String) IndexByte(val byte) int {
	return strings.IndexByte(string(s), val)
}

func (s String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(string(s), f)
}

func (s String) IndexRune(r rune) int {
	return strings.IndexRune(string(s), r)
}

func (s String) LastIndex(subStr String) int {
	return strings.LastIndex(string(s), string(subStr))
}

func (s String) LastIndexAny(chars String) int {
	return strings.LastIndexAny(string(s), string(chars))
}

func (s String) LastIndexByte(val byte) int {
	return strings.LastIndexByte(string(s), val)
}

func (s String) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(string(s), f)
}

func (s String) Map(mapping func(rune) rune) String {
	return String(strings.Map(mapping, string(s)))
}

func (s String) Repeat(count int) String {
	return String(strings.Repeat(string(s), count))
}

func (s String) Replace(target, replacement String, n int) String {
	return String(strings.Replace(string(s), string(target), string(replacement), n))
}

func (s String) ReplaceAll(target, replacement String) String {
	return String(strings.ReplaceAll(string(s), string(target), string(replacement)))
}

func (s String) Split(sep String) Strings {
	return Slice(strings.Split(string(s), string(sep))...)
}

func (s String) SplitAfter(sep String) Strings {
	return Slice(strings.SplitAfter(string(s), string(sep))...)
}

func (s String) SplitAfterN(sep String, n int) Strings {
	return Slice(strings.SplitAfterN(string(s), string(sep), n)...)
}

func (s String) SplitN(sep String, n int) []String {
	return Slice(strings.SplitN(string(s), string(sep), n)...)
}

func (s String) ToLower() String {
	return String(strings.ToLower(string(s)))
}

func (s String) ToLowerSpecial(c unicode.SpecialCase) String {
	return String(strings.ToLowerSpecial(c, string(s)))
}

func (s String) ToTitle() String {
	return String(strings.ToTitle(string(s)))
}

func (s String) ToTitleSpecial(c unicode.SpecialCase) String {
	return String(strings.ToTitleSpecial(c, string(s)))
}

func (s String) ToUpper() String {
	return String(strings.ToUpper(string(s)))
}

func (s String) ToUpperSpecial(c unicode.SpecialCase) String {
	return String(strings.ToUpperSpecial(c, string(s)))
}

func (s String) ToValidUTF8(replacement String) String {
	return String(strings.ToValidUTF8(string(s), string(replacement)))
}

func (s String) Trim(cutset String) String {
	return String(strings.Trim(string(s), string(cutset)))
}

func (s String) TrimFunc(f func(rune) bool) String {
	return String(strings.TrimFunc(string(s), f))
}

func (s String) TrimLeft(cutset String) String {
	return String(strings.TrimLeft(string(s), string(cutset)))
}

func (s String) TrimLeftFunc(f func(rune) bool) String {
	return String(strings.TrimLeftFunc(string(s), f))
}

func (s String) TrimPrefix(prefix String) String {
	return String(strings.TrimPrefix(string(s), string(prefix)))
}

func (s String) TrimRight(cutset String) String {
	return String(strings.TrimRight(string(s), string(cutset)))
}

func (s String) TrimRightFunc(f func(rune) bool) String {
	return String(strings.TrimRightFunc(string(s), f))
}

func (s String) TrimSpace() String {
	return String(strings.TrimSpace(string(s)))
}

func (s String) TrimSuffix(suffix String) String {
	return String(strings.TrimSuffix(string(s), string(suffix)))
}
