package str

import (
	"strings"
	"unicode"
)

//Clone wraps strings.Clone
func (s Str) Clone() Str {
	return Str(strings.Clone(string(s)))
}

//Compare wraps strings.Compare
func (s Str) Compare(other Str) int {
	return strings.Compare(string(s), string(other))
}

//Contains wraps strings.Contains
func (s Str) Contains(subStr Str) bool {
	return strings.Contains(string(s), string(subStr))
}

//ContainsAny wraps ContainsAny
func (s Str) ContainsAny(chars Str) bool {
	return strings.ContainsAny(string(s), string(chars))
}

//ContainsRune wraps strings.ContainsRune
func (s Str) ContainsRune(r rune) bool {
	return strings.ContainsRune(string(s), r)
}

//Count wraps strings.Count
func (s Str) Count(subStr Str) int {
	return strings.Count(string(s), string(subStr))
}

//Cut wraps strings.Cut
func (s Str) Cut(sep Str) (before, after Str, found bool) {
	prefix, suffix, found := strings.Cut(string(s), string(sep))
	return Str(prefix), Str(suffix), found
}

//EqualFold wraps strings.EqualFold
func (s Str) EqualFold(t Str) bool {
	return strings.EqualFold(string(s), string(t))
}

//Fields wraps strings.Fields
func (s Str) Fields() Strs {
	return Slice(strings.Fields(string(s))...)
}

//FieldsFunc wraps strings.FieldsFunc
func (s Str) FieldsFunc(f func(rune) bool) Strs {
	return Slice(strings.FieldsFunc(string(s), f)...)
}

//HasPrefix wraps strings.HasPrefix
func (s Str) HasPrefix(prefix Str) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

//HasSuffix wraps strings.HasSuffix
func (s Str) HasSuffix(suffix Str) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

//Index wraps strings.Index
func (s Str) Index(subStr Str) int {
	return strings.Index(string(s), string(subStr))
}

//IndexAny wraps strings.IndexAny
func (s Str) IndexAny(chars Str) int {
	return strings.IndexAny(string(s), string(chars))
}

//IndexByte wraps strings.IndexByte
func (s Str) IndexByte(val byte) int {
	return strings.IndexByte(string(s), val)
}

//IndexFunc wraps strings.IndexFunc
func (s Str) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(string(s), f)
}

//IndexRune wraps strings.IndexRune
func (s Str) IndexRune(r rune) int {
	return strings.IndexRune(string(s), r)
}

//LastIndex wraps strings.LastIndex
func (s Str) LastIndex(subStr Str) int {
	return strings.LastIndex(string(s), string(subStr))
}

//LastIndexAny wraps strings.LastIndexAny
func (s Str) LastIndexAny(chars Str) int {
	return strings.LastIndexAny(string(s), string(chars))
}

//LastIndexByte wraps strings.LastIndexByte
func (s Str) LastIndexByte(val byte) int {
	return strings.LastIndexByte(string(s), val)
}

//LastIndexFunc wraps strings.LastIndexFunc
func (s Str) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(string(s), f)
}

//Map wraps strings.Map
func (s Str) Map(mapping func(rune) rune) Str {
	return Str(strings.Map(mapping, string(s)))
}

//Repeat wraps strings.Repeat
func (s Str) Repeat(count int) Str {
	return Str(strings.Repeat(string(s), count))
}

//Replace wraps strings.Replace
func (s Str) Replace(target, replacement Str, n int) Str {
	return Str(strings.Replace(string(s), string(target), string(replacement), n))
}

//ReplaceAll wraps strings.ReplaceAll
func (s Str) ReplaceAll(target, replacement Str) Str {
	return Str(strings.ReplaceAll(string(s), string(target), string(replacement)))
}

//Split wraps strings.Split
func (s Str) Split(sep Str) Strs {
	return Slice(strings.Split(string(s), string(sep))...)
}

//SplitAfter wraps strings.SplitAfter
func (s Str) SplitAfter(sep Str) Strs {
	return Slice(strings.SplitAfter(string(s), string(sep))...)
}

//SplitAfterN wraps strings.SplitAfterN
func (s Str) SplitAfterN(sep Str, n int) Strs {
	return Slice(strings.SplitAfterN(string(s), string(sep), n)...)
}

//SplitN wraps strings.SplitN
func (s Str) SplitN(sep Str, n int) []Str {
	return Slice(strings.SplitN(string(s), string(sep), n)...)
}

//ToLower wraps strings.ToLower
func (s Str) ToLower() Str {
	return Str(strings.ToLower(string(s)))
}

//ToLowerSpecial wraps strings.ToLowerSpecial
func (s Str) ToLowerSpecial(c unicode.SpecialCase) Str {
	return Str(strings.ToLowerSpecial(c, string(s)))
}

//ToTitle wraps strings.ToTitle
func (s Str) ToTitle() Str {
	return Str(strings.ToTitle(string(s)))
}

//ToTitleSpecial wraps strings.ToTitleSpecial
func (s Str) ToTitleSpecial(c unicode.SpecialCase) Str {
	return Str(strings.ToTitleSpecial(c, string(s)))
}

//ToUpper wraps strings.ToUpper
func (s Str) ToUpper() Str {
	return Str(strings.ToUpper(string(s)))
}

//ToUpperSpecial wraps strings.ToUpperSpecial
func (s Str) ToUpperSpecial(c unicode.SpecialCase) Str {
	return Str(strings.ToUpperSpecial(c, string(s)))
}

//ToValidUTF8 wraps strings.ToValidUTF8
func (s Str) ToValidUTF8(replacement Str) Str {
	return Str(strings.ToValidUTF8(string(s), string(replacement)))
}

//Trim wraps strings.Trim
func (s Str) Trim(cutset Str) Str {
	return Str(strings.Trim(string(s), string(cutset)))
}

//TrimFunc wraps strings.TrimFunc
func (s Str) TrimFunc(f func(rune) bool) Str {
	return Str(strings.TrimFunc(string(s), f))
}

//TrimLeft wraps strings.TrimLeft
func (s Str) TrimLeft(cutset Str) Str {
	return Str(strings.TrimLeft(string(s), string(cutset)))
}

//TrimLeftFunc wraps strings.TrimLeftFunc
func (s Str) TrimLeftFunc(f func(rune) bool) Str {
	return Str(strings.TrimLeftFunc(string(s), f))
}

//TrimPrefix wraps strings.TrimPrefix
func (s Str) TrimPrefix(prefix Str) Str {
	return Str(strings.TrimPrefix(string(s), string(prefix)))
}

//TrimRight wraps strings.TrimRight
func (s Str) TrimRight(cutset Str) Str {
	return Str(strings.TrimRight(string(s), string(cutset)))
}

//TrimRightFunc wraps strings.TrimRightFunc
func (s Str) TrimRightFunc(f func(rune) bool) Str {
	return Str(strings.TrimRightFunc(string(s), f))
}

//TrimSpace wraps strings.TrimSpace
func (s Str) TrimSpace() Str {
	return Str(strings.TrimSpace(string(s)))
}

//TrimSuffix wraps strings.TrimSuffix
func (s Str) TrimSuffix(suffix Str) Str {
	return Str(strings.TrimSuffix(string(s), string(suffix)))
}
