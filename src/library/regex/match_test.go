package regex

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Printf("%#v\n", validID)
	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]")) //包含大写字母
	fmt.Println(validID.MatchString("snakey"))  // 缺少数字

	// 返回原正则表达式的字符串。
	fmt.Printf("%#v\n", validID.String())
	/*
	 * n> 0：最多n个子字符串； 最后一个子字符串将是未拆分的余数。
	 * n == 0：结果为nil（零子字符串）
	 * n <0：所有子字符串
	 */
	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
	fmt.Printf("%#v\n", s)

	// Match 用于找出字节切片中是否包含该正则表达式的任何匹配项。
	s1 := regexp.MustCompile("a+").Match([]byte("abaabaccadaaae"))
	fmt.Printf("%#v\n", s1)
	s1 = regexp.MustCompile("a+").Match([]byte{'b', 'c'}) // 不包含 a
	fmt.Printf("%#v\n", s1)

}

func TestMatchReader(t *testing.T) {
	// MatchReader 找出 RuneReader 返回的文本是否包含正则表达式的任何匹配项。
	bf0 := bytes.NewBuffer([]byte("大家好，我们来自北京"))
	bf1 := bytes.NewBuffer([]byte("hello, 你们好"))
	s0 := regexp.MustCompile("我们").MatchReader(bf0)
	fmt.Printf("%#v\n", s0)
	s1 := regexp.MustCompile("我们").MatchReader(bf1)
	fmt.Printf("%#v\n", s1)
}

func TestFind(t *testing.T) {
	//Find返回一个切片，其中包含正则表达式b中最左边匹配的文本。
	//返回值nil表示匹配不成功。
	by1 := regexp.MustCompile("a+").Find([]byte("abaabaccadaaae"))
	fmt.Printf("%v\n", by1)
	by1 = regexp.MustCompile("a+").Find([]byte("bc"))
	fmt.Printf("%v\n", by1)

	// FindAll是Find的“all”版本； 它返回表达式的所有连续匹配的[]byte切片，如程序包注释中的“all”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by2 := regexp.MustCompile("a+").FindAll([]byte("abaabaccadaaae"), 0)
	fmt.Printf("%v\n", by2)
	by2 = regexp.MustCompile("a+").FindAll([]byte("abaabaccadaaae"), 1)
	fmt.Printf("%v\n", by2)
	by2 = regexp.MustCompile("a+").FindAll([]byte("abaabaccadaaae"), 2)
	fmt.Printf("%v\n", by2)
	by2 = regexp.MustCompile("a+").FindAll([]byte("bc"), 2)
	fmt.Printf("%v\n", by2)

	// FindAllString是FindString的“全部”版本； 它返回表达式的所有连续匹配的字符串切片，如程序包注释中的“全部”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by33 := regexp.MustCompile("a+").FindAllString("abaabaccadaaae", 0)
	fmt.Printf("%v\n", by33)
	by33 = regexp.MustCompile("a+").FindAllString("abaabaccadaaae", 1)
	fmt.Printf("%v\n", by33)
	by33 = regexp.MustCompile("a+").FindAllString("abaabaccadaaae", 2)
	fmt.Printf("%v\n", by33)
	by33 = regexp.MustCompile("a+").FindAllString("bc", 2)
	fmt.Printf("%v\n", by33)
}

func TestFindIndex(t *testing.T) {
	// FindIndex返回一个由两个元素组成的整数切片，该切片定义正则表达式b中最左边的匹配项的位置。 匹配项本身位于b[loc[0]：loc[1]](不包括loc[1])。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by3 := regexp.MustCompile("a+").FindIndex([]byte("abaabaccadaaae"))
	fmt.Printf("%v\n", by3)
	by3 = regexp.MustCompile("a+").FindIndex([]byte("bc"))
	fmt.Printf("%v\n", by3)

	// FindAll是Find的“all”版本； 它返回表达式的所有连续匹配的一部分，如程序包注释中的“all”描述所定义。
	//返回值nil表示不匹配。
	//如果'All'出现了，该方法会返回输入中所有互不重叠的匹配结果。如果一个匹配结果的前后（没有间隔字符）存在长度为0的成功匹配，
	// 该空匹配会被忽略。包含All的方法会要求一个额外的整数参数n，如果n>=0，方法会返回最多前n个匹配结果。
	by31 := regexp.MustCompile("a+").FindAllIndex([]byte("abaabaccadaaae"), 0)
	fmt.Printf("%v\n", by31)
	by31 = regexp.MustCompile("a+").FindAllIndex([]byte("abaabaccadaaae"), 1)
	fmt.Printf("%v\n", by31)
	by31 = regexp.MustCompile("a+").FindAllIndex([]byte("abaabaccadaaae"), 2)
	fmt.Printf("%v\n", by31)
	by31 = regexp.MustCompile("a+").FindAllIndex([]byte("bc"), 2)
	fmt.Printf("%v\n", by31)
}
