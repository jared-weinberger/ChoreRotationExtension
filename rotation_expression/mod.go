package rotationexpression

import (
	"errors"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
	"github.com/jared-weinberger/ChoreRotationExtension/rotation_expression/parsing"
	"github.com/rs/zerolog/log"
)

type appRotationExpressionListener struct {
	*parsing.BaseRotationExpressionListener
}

func newAppRotationExpressionListener() *appRotationExpressionListener {
	return new(appRotationExpressionListener)
}

func (listener *appRotationExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	log.Info().Msg(ctx.GetText())
}

func Foo(taskBody string) error {
	pattern := regexp.MustCompile("(?i)rotates")
	expressionBeginnings := pattern.FindAllStringIndex(taskBody, -1)
	if expressionBeginnings == nil {
		return errors.New("No rotation expression found!")
	}
	for start, _ := range expressionBeginnings {
		input := antlr.NewInputStream(taskBody[start:])
		lexer := parsing.NewRotationExpressionLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		expression_parser := parsing.NewRotationExpressionParser(stream)
		expression_parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		expression_parser.BuildParseTrees = true
		tree := expression_parser.RotationExpression()
		antlr.ParseTreeWalkerDefault.Walk(newAppRotationExpressionListener(), tree)
	}
	return nil
}
