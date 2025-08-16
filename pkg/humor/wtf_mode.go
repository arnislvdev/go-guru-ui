package humor

import (
	"fmt"
	"math/rand"
	"time"
)

// WTFMode adds humor, sarcasm, and memes to error explanations
type WTFMode struct {
	responses map[string][]string
}

// NewWTFMode creates a new WTF mode instance
func NewWTFMode() *WTFMode {
	wtf := &WTFMode{
		responses: make(map[string][]string),
	}

	wtf.initializeResponses()
	return wtf
}

// Enhance adds humor to an explanation based on error type
func (w *WTFMode) Enhance(explanation, errorType string) string {
	rand.Seed(time.Now().UnixNano())

	// Add a humorous prefix
	prefix := w.getRandomPrefix(errorType)

	// Add a humorous suffix
	suffix := w.getRandomSuffix(errorType)

	// Combine everything
	result := fmt.Sprintf("%s\n\n%s\n\n%s", prefix, explanation, suffix)

	return result
}

// getRandomPrefix returns a random humorous prefix
func (w *WTFMode) getRandomPrefix(errorType string) string {
	prefixes := w.responses["prefixes"]
	if len(prefixes) == 0 {
		return "🤦‍♂️ Oh boy, here we go again..."
	}

	return prefixes[rand.Intn(len(prefixes))]
}

// getRandomSuffix returns a random humorous suffix
func (w *WTFMode) getRandomSuffix(errorType string) string {
	// Get type-specific responses if available
	if responses, exists := w.responses[errorType]; exists && len(responses) > 0 {
		return responses[rand.Intn(len(responses))]
	}

	// Fall back to general responses
	general := w.responses["general"]
	if len(general) == 0 {
		return "💡 Pro tip: Google is your friend... until it's not."
	}

	return general[rand.Intn(len(general))]
}

// initializeResponses sets up the humor responses
func (w *WTFMode) initializeResponses() {
	w.responses["prefixes"] = []string{
		"🤦‍♂️ Oh boy, here we go again...",
		"😅 Well, well, well... what do we have here?",
		"🤔 Let me put on my debugging glasses...",
		"😤 Another day, another error to explain...",
		"🎭 Time for some error theater!",
	}

	w.responses["undefined_symbol"] = []string{
		"🎯 Pro tip: If you can't see it, it probably doesn't exist. Schrödinger's variable!",
		"🔍 Remember: imports are like bringing friends to a party - you have to invite them first!",
		"💡 This is like trying to use a word that's not in your vocabulary. Time to expand your horizons!",
	}

	w.responses["type_mismatch"] = []string{
		"🔄 It's like trying to fit a square peg in a round hole. Spoiler: it won't work!",
		"🎭 Type casting is like acting - sometimes you need to change your role!",
		"🔧 Remember: Go is strongly typed, not strongly opinionated about your choices!",
	}

	w.responses["unused_import"] = []string{
		"🗑️ You're like that person who brings a suitcase to a day trip - overpacked!",
		"🧹 Clean code is happy code. Time to Marie Kondo your imports!",
		"📦 Import what you need, not what you might need someday!",
	}

	w.responses["unused_variable"] = []string{
		"👻 You've created a ghost variable - it exists but nobody sees it!",
		"🕵️ This variable is so stealthy, even your code forgot about it!",
		"🎪 You're running a one-variable show with no audience!",
	}

	w.responses["missing_return"] = []string{
		"🚪 You promised to return something but forgot to actually leave the building!",
		"📤 Functions are like promises - you gotta keep them!",
		"🏃‍♂️ You started the race but never crossed the finish line!",
	}

	w.responses["argument_count"] = []string{
		"🎯 You're either overfeeding or starving your function!",
		"📊 Function arguments are like Goldilocks - not too many, not too few, just right!",
		"🔢 Math is hard, but counting arguments shouldn't be!",
	}

	w.responses["general"] = []string{
		"💡 Pro tip: Google is your friend... until it's not.",
		"🎭 Remember: every error is just a learning opportunity in disguise!",
		"🚀 You're one step closer to being a debugging master!",
		"🔧 Keep calm and debug on!",
		"🎪 Welcome to the wonderful world of programming errors!",
	}
}
