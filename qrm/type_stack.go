package qrm

import "reflect"

type typeStack []*reflect.Type

func newTypeStack() typeStack { _ = "STUB: not implemented"; return *new(typeStack) }

func (s *typeStack) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *typeStack) push(t *reflect.Type) { _ = "STUB: not implemented"; return }

func (s *typeStack) pop() bool { _ = "STUB: not implemented"; return false }

func (s *typeStack) contains(t *reflect.Type) bool { _ = "STUB: not implemented"; return false }
