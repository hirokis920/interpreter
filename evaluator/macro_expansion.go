package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}

	for i, statement := range program.Statements {
		if isMacroDifinision(statement) {
			addMacro(statement, env)
			definitions = append(definitions, i)
		}
	}
	//todo ここのコロンの意味を調べる
	for i := len(definitions) - 1; i >= 0; i = i - 1 {
		definitionIndex := definitions[i]
		program.Statements = append(program.Statements[:definitionIndex], program.Statements[definitionIndex+1:]...)
	}
}
