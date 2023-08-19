package main

import "fmt"

// Condition represents a condition that needs to be satisfied.
type Condition func(factors map[string]interface{}) bool

// Action represents an action to be taken when conditions are met.
type Action func(factors map[string]interface{})

// Rule represents a rule with its conditions and action.
type Rule struct {
	Conditions []Condition
	Action     Action
}

// RuleEngine represents the rule engine.
type RuleEngine struct {
	Rules []Rule
}

// NewRuleEngine creates a new RuleEngine instance.
func NewRuleEngine() *RuleEngine {
	return &RuleEngine{}
}

// AddRule adds a rule to the RuleEngine.
func (re *RuleEngine) AddRule(conditions []Condition, action Action) {
	re.Rules = append(re.Rules, Rule{Conditions: conditions, Action: action})
}

// EvaluateRules evaluates the rules against the given factors.
func (re *RuleEngine) EvaluateRules(factors map[string]interface{}) {
	for _, rule := range re.Rules {
		conditionsMet := true
		for _, cond := range rule.Conditions {
			if !cond(factors) {
				conditionsMet = false
				break
			}
		}
		if conditionsMet {
			rule.Action(factors)
		}
	}
}

func main() {
	ruleEngine := NewRuleEngine()

	ruleEngine.AddRule(
		[]Condition{
			func(factors map[string]interface{}) bool {
				temperature, ok := factors["temperature"].(float64)
				return ok && temperature > 30
			},
		},
		func(factors map[string]interface{}) {
			fmt.Println("It's hot! Turn on the AC.")
		},
	)

	factors := map[string]interface{}{
		"temperature": 32.0,
	}
	ruleEngine.EvaluateRules(factors)
}
