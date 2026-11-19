package errregistry

import (
	"errors"
	"net/http"
)

type rule struct {
	match func(error) (APIError, bool)
}

type Registry struct {
	rules       []rule
	fallbackErr APIError
}

func New() *Registry {
	return &Registry{
		fallbackErr: APIError{
			HTTPStatus: http.StatusInternalServerError,
			Code:       "INTERNAL_ERROR",
			Message:    "Unexpected server error",
		},
	}
}

func (r *Registry) RegisterSentinel(target error, resp APIError) {
	r.rules = append(r.rules, rule{
		match: func(err error) (APIError, bool) {
			if errors.Is(err, target) {
				return resp, true
			}
			return APIError{}, false
		},
	})
}

func RegisterDynamic[T error](r *Registry, mapper func(T) APIError) {
	r.rules = append(r.rules, rule{
		match: func(err error) (APIError, bool) {
			var target T
			if errors.As(err, &target) {
				return mapper(target), true
			}
			return APIError{}, false
		},
	})
}

func (r *Registry) Merge(other *Registry) {
	if other == nil {
		return
	}
	r.rules = append(r.rules, other.rules...)
}

func (r *Registry) Resolve(err error) APIError {
	for _, rule := range r.rules {
		if resp, ok := rule.match(err); ok {
			return resp
		}
	}
	return r.fallbackErr
}
