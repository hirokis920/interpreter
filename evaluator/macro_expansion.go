package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// todo ネストしたマクロに対応する
func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}

	for i, statement := range program.Statements {
		//マクロ定義かどうかを判定

		if isMacroDifinision(statement) {
			//後で削除できるように位置を記録
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

func isMacroDifinision(node ast.Statement) bool {
	letStatement, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStatement.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}

	return true
}

func addMacro(stmt ast.Statement, env *object.Environment) {
	letStatement, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStatement.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        env,
		Body:       macroLiteral.Body,
	}
	env.Set(letStatement.Name.Value, macro)
}
