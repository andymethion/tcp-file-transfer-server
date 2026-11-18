package proto

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

func GetCommand(request string) (command string) {
	return strings.Split(request, " ")[0]
}

func GetArgument(request string) (argument string) {
	splitN := strings.SplitN(request, " ", 2)

	if len(splitN) != 2 {
		return ""
	}

	return splitN[1]
}

func Write(timeout time.Duration, conn net.Conn, s string) (err error) {
	conn.SetDeadline(time.Now().Add(timeout))
	_, err = conn.Write([]byte(s))
	conn.SetDeadline(time.Time{})
	return err
}

func Read(timeout time.Duration, conn net.Conn, b []byte) (s string, err error) {
	conn.SetDeadline(time.Now().Add(timeout))
	n, err := conn.Read(b)
	conn.SetDeadline(time.Time{})
	return string(b[:n]), err
}

func Input(prompt string) (answer string, err error) {
	fmt.Printf(prompt)

	reader := bufio.NewReader(os.Stdin)
	answer, err = reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(answer), nil
}

// Affiche un log avec le client et son erreur
func SlogError(err error) {
	typeOfErr := reflect.TypeOf(err)
	_, file, line, _ := runtime.Caller(1)
	slog.Error(err.Error(), "typeOfErr", typeOfErr, "file", file, "line", line)
}

// Renvoie le nombre d'élément d'une sync.Map (le nombre de clients)
func CountMap(clients *sync.Map) (count int) {
	count = 0

	clients.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	return count
}
