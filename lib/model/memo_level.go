package model

import (
  "os"
  "fmt"
  "regexp"

  "gopkg.in/yaml.v2"
)

// MemoLevel is
type MemoLevel struct{}

type LevelYaml struct {
  Level1 []string `yaml:"level1"`
  Level2 []string `yaml:"level2"`
  Level3 []string `yaml:"level3"`
}

func GetLevelYaml() (LevelYaml, error) {
  levelYaml := LevelYaml{}
  b, err := os.ReadFile("lib/mode/level.yml")

  if err != nil {
    return LevelYaml{}, err
  }
  err = yaml.Unmarshal(b, &levelYaml)

  if err != nil {
    return LevelYaml{}, err
  }

  return levelYaml, nil
}

func getSharp(level int) string {
  sharp := ""
  for index := 0; index < level; index++ {
    sharp += "#"
  }

  return sharp
}

func getLevel(line string, level int) (string, string, error) {
  sharp := getSharp(level)
  re := regexp.MustCompile(fmt.Sprintf(`^%s\s\[.+\]\s(.*)`, sharp))
  levelYamls , err := GetLevelYaml()
  if err != nil {
    return "", "", err
  }

  levelYaml := []string{}
  if level == 3 {
    levelYaml = levelYamls.Level3
  } else if level == 2 {
    levelYaml = levelYamls.Level2
  } else if level == 1 {
    levelYaml = levelYamls.Level1
  } else {
    return "", "", fmt.Errorf(fmt.Sprintf("Dose not exist level%d", level))
  }

  for index := range levelYaml {
    text := levelYaml[index]
    matchText := fmt.Sprintf(`^%s\s\[%s\]`, sharp, text)
    // fmt.Println(line)
    // fmt.Println(fmt.Sprintf(`^%s\s\[%s\]`, sharp, text))
    // fmt.Println(regexp.MustCompile(matchText).Match([]byte(line)))
    if regexp.MustCompile(matchText).Match([]byte(line)) {
      return re.FindStringSubmatch(line)[1], text, nil
    }
  }

  return "", "", fmt.Errorf(fmt.Sprintf("Not found level text in level%d", level))
}

// GetLevel3 is
func getLevel3(line string) (string, string) {
  title, text, _ := getLevel(line, 3)
  return title, text
}

// GetLevel2 is
func getLevel2(line string) (string, string) {
  title, text, _ := getLevel(line, 2)
  return title, text
}

// GetLevel1 is
func getLevel1(line string) (string, string) {
  title, text, _ := getLevel(line, 1)
  return title, text
}
