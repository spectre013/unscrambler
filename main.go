package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	b := map[int]string{}
	a := buildCharArray()
	word := "test"
	for k := range word {
		b[k] = string(word[k])
		delete(a, string(word[k]))
	}

	e := "cat combined.txt"
	//out, err := exec.Command("cat"," combined.txt | grep -v $k").Output()
	for k := range a {
		//fmt.Println(k)
		e += " | grep -v " + k
	}
	//log.Println(e)
	out, err := exec.Command("bash", "-c", e).Output()
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(e)
	fmt.Printf("%s\n", out)

}

func buildCharArray() map[string]int {
	var a = map[string]int{}
	for i := 65; i <= 90; i++ {
		str := toCharStr(i)
		a[str] = 1
	}
	return a
}

func toCharStr(i int) string {
	return strings.ToLower(string(i))
}

func contains(word string) bool {
	wlen := len(word)
	if wlen < 3 || wlen > 10 {
		return false
	}

	return true
}

/*
function contains($word) {
    $wlen = strlen($word);
    if ($wlen < 3 || $wlen > 10) {
        return false;
    }

    for ($i = 0; $i < $wlen; $i++) {
        $w[$i] = $word[$i];
    }

    $b = $GLOBALS['b'];
    foreach ($b as $n => $c) {
        foreach ($w as $k => $v) {
            if ($v == $c) {
                unset($w[$k]);
                unset($b[$n]);
                break;
            }
        }
    }
    //echo "count: ". count($w) ."\n";
    if (count($w) > 0) {
        return false;
    }

    return true;
}
*/
