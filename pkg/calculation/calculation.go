package calculation

import (
	"errors"
	"strconv"
	"unicode"
)

// Функция для вычисления арифметического выражения, используя алгоритм сортировочной станции.
func Calc(expression string) (float64, error) {
	var output []string // Постфиксное выражение (RPN)
	var ops []rune      // Стек операторов
	var num string      // Для накопления чисел

	// Приоритеты операторов
	priority := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	// Преобразование выражения в обратную польскую нотацию (RPN)
	for i, ch := range expression {
		switch {
		case unicode.IsDigit(ch) || ch == '.': // Числа
			num += string(ch)
			// Если это последний символ числа, добавляем его в выходную строку
			if i == len(expression)-1 || (!unicode.IsDigit(rune(expression[i+1])) && expression[i+1] != '.') {
				output = append(output, num)
				num = ""
			}
		case ch == '+' || ch == '-' || ch == '*' || ch == '/': // Операторы
			for len(ops) > 0 && ops[len(ops)-1] != '(' && priority[ops[len(ops)-1]] >= priority[ch] {
				output = append(output, string(ops[len(ops)-1]))
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, ch)
		case ch == '(': // Открывающая скобка
			ops = append(ops, ch)
		case ch == ')': // Закрывающая скобка
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				output = append(output, string(ops[len(ops)-1]))
				ops = ops[:len(ops)-1]
			}
			if len(ops) == 0 || ops[len(ops)-1] != '(' {
				return 0, errors.New("mismatched parentheses")
			}
			ops = ops[:len(ops)-1] // Убираем '(' из стека
		case unicode.IsSpace(ch): // Игнорируем пробелы
			continue
		default:
			return 0, errors.New("invalid character in expression")
		}
	}

	// Выгружаем оставшиеся операторы из стека
	for len(ops) > 0 {
		if ops[len(ops)-1] == '(' || ops[len(ops)-1] == ')' {
			return 0, errors.New("mismatched parentheses")
		}
		output = append(output, string(ops[len(ops)-1]))
		ops = ops[:len(ops)-1]
	}

	// Вычисление значения выражения в формате RPN
	var evalStack []float64
	for _, token := range output {
		switch token {
		case "+", "-", "*", "/":
			if len(evalStack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := evalStack[len(evalStack)-1], evalStack[len(evalStack)-2]
			evalStack = evalStack[:len(evalStack)-2]
			var res float64
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				res = a / b
			}
			evalStack = append(evalStack, res)
		default: // Числа
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("invalid number")
			}
			evalStack = append(evalStack, num)
		}
	}

	if len(evalStack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return evalStack[0], nil
}
