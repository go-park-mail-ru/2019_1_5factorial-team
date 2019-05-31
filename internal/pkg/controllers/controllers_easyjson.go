// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package controllers

import (
	json "encoding/json"
	user "github.com/go-park-mail-ru/2019_1_5factorial-team/internal/pkg/user"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers(in *jlexer.Lexer, out *signInRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "loginOrEmail":
			out.LoginOrEmail = string(in.String())
		case "password":
			out.Password = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers(out *jwriter.Writer, in signInRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"loginOrEmail\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LoginOrEmail))
	}
	{
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v signInRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v signInRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *signInRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *signInRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers(l, v)
}
func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers1(in *jlexer.Lexer, out *UsersCountInfoResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "count":
			out.Count = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers1(out *jwriter.Writer, in UsersCountInfoResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Count))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersCountInfoResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersCountInfoResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersCountInfoResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersCountInfoResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers1(l, v)
}
func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers2(in *jlexer.Lexer, out *UserInfoResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "email":
			out.Email = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "score":
			out.Score = int(in.Int())
		case "avatar_link":
			out.AvatarLink = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers2(out *jwriter.Writer, in UserInfoResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"score\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Score))
	}
	{
		const prefix string = ",\"avatar_link\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarLink))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInfoResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInfoResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInfoResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInfoResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers2(l, v)
}
func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers3(in *jlexer.Lexer, out *SingUpRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "login":
			out.Login = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "avatar_link":
			out.AvatarLink = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers3(out *jwriter.Writer, in SingUpRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"avatar_link\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarLink))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SingUpRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SingUpRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SingUpRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SingUpRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers3(l, v)
}
func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers4(in *jlexer.Lexer, out *ProfileUpdateResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "email":
			out.Email = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "score":
			out.Score = int(in.Int())
		case "avatar_link":
			out.AvatarLink = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers4(out *jwriter.Writer, in ProfileUpdateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"score\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Score))
	}
	{
		const prefix string = ",\"avatar_link\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarLink))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProfileUpdateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProfileUpdateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProfileUpdateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers4(l, v)
}
func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers5(in *jlexer.Lexer, out *ProfileUpdateRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "avatar":
			out.Avatar = string(in.String())
		case "old_password":
			out.OldPassword = string(in.String())
		case "new_password":
			out.NewPassword = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers5(out *jwriter.Writer, in ProfileUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"avatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"old_password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OldPassword))
	}
	{
		const prefix string = ",\"new_password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NewPassword))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProfileUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProfileUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProfileUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProfileUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers5(l, v)
}
func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers6(in *jlexer.Lexer, out *GetLeaderboardResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "scores":
			if in.IsNull() {
				in.Skip()
				out.Scores = nil
			} else {
				in.Delim('[')
				if out.Scores == nil {
					if !in.IsDelim(']') {
						out.Scores = make([]user.Scores, 0, 2)
					} else {
						out.Scores = []user.Scores{}
					}
				} else {
					out.Scores = (out.Scores)[:0]
				}
				for !in.IsDelim(']') {
					var v1 user.Scores
					(v1).UnmarshalEasyJSON(in)
					out.Scores = append(out.Scores, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers6(out *jwriter.Writer, in GetLeaderboardResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"scores\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Scores == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Scores {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetLeaderboardResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetLeaderboardResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetLeaderboardResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetLeaderboardResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers6(l, v)
}
func easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers7(in *jlexer.Lexer, out *AvatarLinkResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "avatar_link":
			out.AvatarLink = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers7(out *jwriter.Writer, in AvatarLinkResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"avatar_link\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarLink))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AvatarLinkResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AvatarLinkResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson650569b1EncodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AvatarLinkResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AvatarLinkResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson650569b1DecodeGithubComGoParkMailRu201915factorialTeamInternalPkgControllers7(l, v)
}