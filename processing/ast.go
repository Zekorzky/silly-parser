package processing

type ASTNodeType int

const (
	UNKNOWN ASTNodeType = iota
	NUMBER
	OPERATOR
	FUNCTION
)

type ASTNode struct {
	Type     ASTNodeType
	Value    string
	Children []*ASTNode
}

func CreateAST(input []string) *ASTNode {
	p := Parser{input, 0}
	return p.parsePlusMinus()
}

type Parser struct {
	tokens []string
	pos    int
}

func (p *Parser) peek() string {
	if p.pos >= len(p.tokens) {
		return ""
	}
	return p.tokens[p.pos]
}

func (p *Parser) consume() string {
	op := p.peek()
	p.pos = p.pos + 1
	return op
}

func (p *Parser) parsePlusMinus() *ASTNode {
	left := p.parseMultiDiv()
	for p.peek() == "+" || p.peek() == "-" {
		op := p.consume()
		right := p.parseMultiDiv()
		left = &ASTNode{Type: OPERATOR, Value: op, Children: []*ASTNode{left, right}}
	}
	return left
}

func (p *Parser) parseMultiDiv() *ASTNode {
	left := p.parseExponent()
	for p.peek() == "*" || p.peek() == "/" {
		op := p.consume()
		right := p.parseExponent()
		left = &ASTNode{Type: OPERATOR, Value: op, Children: []*ASTNode{left, right}}
	}
	return left
}

func (p *Parser) parseExponent() *ASTNode {
	base := p.parseFactor()
	for p.peek() == "^" {
		op := p.consume()
		exp := p.parseFactor()
		base = &ASTNode{Type: OPERATOR, Value: op, Children: []*ASTNode{base, exp}}
	}
	return base
}

func (p *Parser) parseFactor() *ASTNode {
	t := p.consume()

	if t == "e" {
		return &ASTNode{Type: NUMBER, Value: "2.718281828459045" /* bro forgot how to convert float to string =( */}
	}
	if t == "pi" {
		return &ASTNode{Type: NUMBER, Value: "3.141592653589793"}
	}
	if t == "(" {
		node := p.parsePlusMinus()
		p.consume()
		return node
	}

	if IsToken(t) {
		p.consume()
		arg := p.parsePlusMinus()
		p.consume()
		return &ASTNode{Type: FUNCTION, Value: t, Children: []*ASTNode{arg}}
	}

	return &ASTNode{Type: NUMBER, Value: t}
}
