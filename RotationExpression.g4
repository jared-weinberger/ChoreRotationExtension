grammar RotationExpression;
options {
	caseInsensitive = true;
}

rotationExpression: PREAMBLE ':'? WHITESPACE userList EOF;
userList: userSpec (',' WHITESPACE userSpec)*;
userSpec: partialUserSpec | unknownUser | completeUserSpec;
unknownUser: UNKNOWN;
partialUserSpec: USERNAME;
completeUserSpec: USERNAME WHITESPACE '(' USER_ID ')';

WHITESPACE: ' '+;
fragment ROTATES: 'rotates';
fragment AMONG: 'among';
PREAMBLE: ROTATES WHITESPACE AMONG;
UNKNOWN: 'unknown';
// USERNAME_SOLID: [a-z0-9._@\-!]+; USERNAME: USERNAME_SOLID+;
USERNAME: 'foo' | 'bar' | 'foobar';
USER_ID: [0-9]+;

// Separator = CommaSeparator | "\n"

// CommaSeparator = "," " "*

// Space = "  " (Space | "")

// OptionalSpace = "" | Space

// AnySpace = " " | "\n" |

// UserRef = "unknown" | PartialUserSpec | CompleteUserSpec

// PartialUserRef = AlphaNum AlphaNum | Space | "." | "-" | "@" | "." | "-" | "@"