package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
	Паттерн Состояние нужно использовать в случаях, когда объект может иметь много различных состояний,
	которые он должен менять в зависимости от конкретного поступившего запроса.
	Плюсы:
	Концентрирует в одном месте код, связанный с определённым состоянием.
	Минусы:
	Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

//Интерфейс состояния
type state interface {
	writingArticle()
	verifyArticle()
	publicationArticle()
}

type newsArticle struct {
	draft      state
	moderation state
	published  state

	currentState state
}

func newNewsArticle() *newsArticle {
	v := &newsArticle{}

	draftState := &draftState{
		newsArticle: v,
	}
	moderationState := &moderationState{
		newsArticle: v,
	}
	publishedState := &publishedState{
		newsArticle: v,
	}
	v.setState(draftState)
	v.draft = draftState
	v.moderation = moderationState
	v.published = publishedState
	return v
}

func (v *newsArticle) writingArticle() {
	v.currentState.writingArticle()
}

func (v *newsArticle) verifyArticle() {
	v.currentState.verifyArticle()
}

func (v *newsArticle) publicationArticle() {
	v.currentState.publicationArticle()
}

func (v *newsArticle) setState(s state) {
	v.currentState = s
}

type draftState struct {
	newsArticle *newsArticle
}

func (v *draftState) writingArticle() {
	v.newsArticle.setState(v.newsArticle.draft)
}

func (v *draftState) verifyArticle() {
	v.newsArticle.setState(v.newsArticle.moderation)
}

func (v *draftState) publicationArticle() {
	fmt.Println("Can't post without moderation")
}

type moderationState struct {
	newsArticle *newsArticle
}

func (v *moderationState) writingArticle() {
	fmt.Println("Article already written")
}

func (v *moderationState) verifyArticle() {
	v.newsArticle.setState(v.newsArticle.moderation)
}

func (v *moderationState) publicationArticle() {
	v.newsArticle.setState(v.newsArticle.published)
}

type publishedState struct {
	newsArticle *newsArticle
}

func (v *publishedState) writingArticle() {
	fmt.Println("Article already written")
}

func (v *publishedState) verifyArticle() {
	fmt.Println("Article already verified")
}

func (v *publishedState) publicationArticle() {
	v.newsArticle.setState(v.newsArticle.published)
}
