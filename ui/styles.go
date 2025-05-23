package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StyleSheet struct {
	States     map[UIState]StyleSet
	Animations map[StyleProperty]AnimationConfig
}

type StyleSet map[StyleProperty]interface{}

type StyleProperty string

const (
	Tint   StyleProperty = "tint"
	Offset StyleProperty = "offset"
	Scale  StyleProperty = "scale"
)

type AnimationConfig struct {
	Duration float64
	Easing   func(t float64) float64
}

type StyleAnimation struct {
	Property StyleProperty
	From     interface{}
	To       interface{}
	Start    float64
	Duration float64
	Easing   func(t float64) float64
}

func ResolveStyle(base *UIBase, prop StyleProperty) interface{} {
	if base == nil || base.Style == nil {
		return nil
	}
	if anim, ok := base.AnimState[prop]; ok {
		t := (currentTime - anim.Start) / anim.Duration
		t = clamp01(anim.Easing(t))
		return interpolate(anim.From, anim.To, t)
	}
	stateSet, ok := base.Style.States[base.State]
	if !ok {
		return nil
	}
	return stateSet[prop]
}

var currentTime float64

func AdvanceAnimations(root Element, dt float32) {
	currentTime += float64(dt)
	advanceRecursive(root, currentTime)
}

func advanceRecursive(el Element, now float64) {
	base := el.GetUIBase()
	for prop, anim := range base.AnimState {
		t := (now - anim.Start) / anim.Duration
		if t >= 1.0 {
			delete(base.AnimState, prop) // Animation complete
		}
	}
	for _, child := range base.Children {
		advanceRecursive(child, now)
	}
}

func SetState(e Element, newState UIState) {
	base := e.GetUIBase()
	old := base.State
	base.State = newState

	if base.Style == nil {
		return
	}

	oldSet := base.Style.States[old]
	newSet := base.Style.States[newState]

	if base.AnimState == nil {
		base.AnimState = make(map[StyleProperty]*StyleAnimation)
	}

	for prop, newVal := range newSet {
		oldVal := oldSet[prop]
		if !styleValuesEqual(oldVal, newVal) {
			cfg := base.Style.Animations[prop]
			base.AnimState[prop] = &StyleAnimation{
				Property: prop,
				From:     oldVal,
				To:       newVal,
				Start:    currentTime,
				Duration: cfg.Duration,
				Easing:   cfg.Easing,
			}
		}
	}
}

func AnimateProperty(base *UIBase, prop StyleProperty, from, to interface{}, duration float64, easing func(float64) float64) {
	if base.AnimState == nil {
		base.AnimState = make(map[StyleProperty]*StyleAnimation)
	}
	now := currentTime
	base.AnimState[prop] = &StyleAnimation{
		Property: prop,
		From:     from,
		To:       to,
		Start:    now,
		Duration: duration,
		Easing:   easing,
	}
}

func interpolate(from, to interface{}, t float64) interface{} {
	switch fromVal := from.(type) {
	case rl.Color:
		toVal := to.(rl.Color)
		return rl.Color{
			R: uint8(lerp(float64(fromVal.R), float64(toVal.R), t)),
			G: uint8(lerp(float64(fromVal.G), float64(toVal.G), t)),
			B: uint8(lerp(float64(fromVal.B), float64(toVal.B), t)),
			A: uint8(lerp(float64(fromVal.A), float64(toVal.A), t)),
		}
	case rl.Vector2:
		toVal := to.(rl.Vector2)
		return rl.NewVector2(
			float32(lerp(float64(fromVal.X), float64(toVal.X), t)),
			float32(lerp(float64(fromVal.Y), float64(toVal.Y), t)),
		)
	case float32:
		toVal := to.(float32)
		return float32(lerp(float64(fromVal), float64(toVal), t))
	default:
		return to
	}
}

func styleValuesEqual(a, b interface{}) bool {
	switch a := a.(type) {
	case rl.Color:
		b := b.(rl.Color)
		return a == b
	case rl.Vector2:
		b := b.(rl.Vector2)
		return a.X == b.X && a.Y == b.Y
	case float32:
		return a == b.(float32)
	default:
		return false
	}
}

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// EASING FUNCTIONS
func EaseOutQuad(t float64) float64 {
	return 1 - (1-t)*(1-t)
}

func EaseInQuad(t float64) float64 {
	return t * t
}

func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return -1 + (4-2*t)*t
}

// DEBUGGING
func PrintActiveAnimations(root Element) {
	printAnimationsRecursive(root)
}

func printAnimationsRecursive(el Element) {
	base := el.GetUIBase()
	if len(base.AnimState) > 0 {
		fmt.Printf("[Animations] ID=%s\n", base.ID)
		for prop, anim := range base.AnimState {
			fmt.Printf("  - Prop: %s | From: %v -> To: %v | Start: %.2f | Duration: %.2f\n",
				prop, anim.From, anim.To, anim.Start, anim.Duration)
		}
	}
	for _, child := range base.Children {
		printAnimationsRecursive(child)
	}
}
