package core

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

// LimitRangeOps is an interface to perform k8s LimitRange operations
type LimitRangeOps interface {
	// GetLimitRange gets the secrets object given its name and namespace
	GetLimitRange(name string, namespace string) (*corev1.LimitRange, error)
	// CreateLimitRange creates the given secret
	CreateLimitRange(*corev1.LimitRange) (*corev1.LimitRange, error)
	// ListLimitRange changes and callback fn
	ListLimitRange(string, metav1.ListOptions) (*corev1.LimitRangeList, error)
	// UpdateLimitRange updates the given secret
	UpdateLimitRange(*corev1.LimitRange) (*corev1.LimitRange, error)
	// DeleteLimitRange deletes the given secret
	DeleteLimitRange(name, namespace string) error
	// WatchLimitRange changes and callback fn
	WatchLimitRange(*corev1.LimitRange, WatchFunc) error
}

// GetLimitRange gets the secrets object given its name and namespace
func (c *Client) GetLimitRange(name string, namespace string) (*corev1.LimitRange, error) {
	if err := c.initClient(); err != nil {
		return nil, err
	}

	return c.kubernetes.CoreV1().LimitRanges(namespace).Get(name, metav1.GetOptions{})
}

// CreateLimitRange creates the given secret
func (c *Client) CreateLimitRange(secret *corev1.LimitRange) (*corev1.LimitRange, error) {
	if err := c.initClient(); err != nil {
		return nil, err
	}

	return c.kubernetes.CoreV1().LimitRanges(secret.Namespace).Create(secret)
}

// ListLimitRange creates the given secret
func (c *Client) ListLimitRange(namespace string, opts metav1.ListOptions) (*corev1.LimitRangeList, error) {
	if err := c.initClient(); err != nil {
		return nil, err
	}

	return c.kubernetes.CoreV1().LimitRanges(namespace).List(opts)
}

// UpdateLimitRange updates the given secret
func (c *Client) UpdateLimitRange(limitrange *corev1.LimitRange) (*corev1.LimitRange, error) {
	if err := c.initClient(); err != nil {
		return nil, err
	}

	return c.kubernetes.CoreV1().LimitRanges(limitrange.Namespace).Update(limitrange)
}

// DeleteLimitRange deletes the given secret
func (c *Client) DeleteLimitRange(name, namespace string) error {
	if err := c.initClient(); err != nil {
		return err
	}

	return c.kubernetes.CoreV1().LimitRanges(namespace).Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deleteForegroundPolicy,
	})
}

func (c *Client) WatchLimitRange(limitrange *v1.LimitRange, fn WatchFunc) error {
	if err := c.initClient(); err != nil {
		return err
	}

	listOptions := metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("metadata.name", limitrange.Name).String(),
		Watch:         true,
	}

	watchInterface, err := c.kubernetes.CoreV1().LimitRanges(limitrange.Namespace).Watch(listOptions)
	if err != nil {
		return err
	}

	// fire off watch function
	go c.handleWatch(watchInterface, limitrange, "", fn, listOptions)
	return nil
}
