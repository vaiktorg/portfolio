package pages

import (
	"github.com/gtank/isaac"
	. "github.com/maxence-charriere/go-app/pkg/app"
	uuid "github.com/satori/go.uuid"
	"sort"
	"time"
)

type GUIDCardShuffler struct {
	Compo
	cards     []Card
	IShuffled bool
}

type Card struct {
	suit  string
	color string
	val   string
}

func (s *GUIDCardShuffler) Render() UI {
	return Div().Body(
		Div().Body(
			H1().Text("GUID SHUFFLING ALGORITHM"),
			P().Text("A small example of generating a safe random card guidShuffle."),
		).Class("w3-center"),

		Hr(),
		Div().Body(
			Pre().Body(Code().Text(`sort.Slice(s.cards, func(i, j int) bool {
		str := uid.String()
		iidx := i % len(str)
		jidx := j % len(str)

		return s.cards[iidx].val < string(str[jidx])
})`)).Class("mycode"),
		),

		Div().Body(
			Input().Class("w3-button w3-margin w3-black").Type("button").Value("GUID Shuffle Deck").OnClick(s.OnClickGUIDShuffle),
			Input().Class("w3-button w3-margin w3-black").Type("button").Value("ISAAC Shuffle Deck").OnClick(s.OnClickISAACShuffle),
			Input().Class("w3-button w3-margin w3-black").Type("button").Value("Sort Deck").OnClick(s.OnClickOrderDeck),
			Br(),
			Range(s.cards).Slice(func(i int) UI {
				return Div().Body(
					H1().Text(s.cards[i].val + s.cards[i].suit).Class("w3-center"),
				).
					Class("w3-col s12 m3 l2 w3-border w3-" + s.cards[i].color)
			}),
		).Class("w3-center"),
	).Class("w3-animate-opacity").ID("content")
}

func (s *GUIDCardShuffler) OnMount(Context) {
	s.populateDeck()
	s.Update()
}

func (s *GUIDCardShuffler) OnClickOrderDeck(Context, Event) {
	s.populateDeck()
	s.Update()
}
func (s *GUIDCardShuffler) OnClickGUIDShuffle(Context, Event) {
	s.guidShuffle()
	s.Update()
}
func (s *GUIDCardShuffler) OnClickISAACShuffle(Context, Event) {
	s.isaacShuffle()
	s.Update()
}

func (s *GUIDCardShuffler) populateDeck() {
	s.cards = []Card{}
	nums := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suitscolor := map[string]string{"♠": "black", "♣": "grey", "♥": "red", "♦": "orange"}

	for suit, color := range suitscolor {
		for _, v := range nums {
			s.cards = append(s.cards, Card{
				suit:  suit,
				color: color,
				val:   v,
			})
		}
	}
}

func (s *GUIDCardShuffler) guidShuffle() {
	uid, err := uuid.NewV4()
	if err != nil {
		return
	}

	sort.Slice(s.cards, func(i, j int) bool {
		str := uid.Bytes()
		jidx := j % len(str)

		return s.cards[i].val < string(str[jidx])
	})

}

func (s *GUIDCardShuffler) isaacShuffle() {
	is := isaac.ISAAC{}
	is.Seed(time.Now().String())
	sort.Slice(s.cards, func(i, _ int) bool {
		idx := int(is.Rand()) % len(s.cards)
		return s.cards[i].val < s.cards[idx].val
	})
}
