# Hangulize 2

(WIP: 아직 개발 중, API가 임의로 바뀔 수 있어요!)

![한글라이즈](brand/logo-kor-128px.png)

> 외국어의 한글 표기 체계가 제대로 서려면 일반인이 외국어를 한글로 표기하고
> 싶을 때 바로바로 쉽게 용례를 찾을 수 있어야 한다. 정기적으로 회의를 열어
> 용례를 정하는 것으로는 한계가 있다. 외래어 표기 심의 방식이 자동화되어 한글로
> 표기하고 싶은 외국어를 입력하자마자 한글 표기가 나와야 한다. 이미 용례가
> 정해진 것은 그것을 따르고 용례에 없는 것이라도 각 언어의 표기 규칙에 따라
> 권장 표기를 표시해야 한다. 프로그래머들과 언어학자들이 손잡고 연구한다면 이게
> 공상으로만 그치지 않을 것이다.
>
> Brian Jongseong Park (http://iceager.egloos.com/2610028)

Hangulize 2는 외래어를 한글로 변환해주는 도구입니다.

```console
$ go get -u github.com/sublee/hangulize2/hangulize
```

```go
import "github.com/sublee/hangulize2/hangulize"

hangulize.Hangulize("ita", "gloria")
// output: "글로리아"
```

## 리부트

Hangulize 2의 전신인 [Hangulize](https://github.com/sublee/hangulize)는
2010년에 Python으로 개발되었습니다. 웹 상에서 사용할 수 있도록
[hangulize.org](http://hangulize.org)를 제공하고 있습니다.

기존 Hangulize의 기능은 모두 계승하면서, 성능을 높이고 코드의 유지보수 가능성과
규칙 설계 시 생산성을 높이기 위해 리부트합니다. Hangulize 2는 Go로 개발됩니다.

## 목표

- 기존 Hangulize 기능 모두 계승
- 규칙 설계에 정적 파일(`.hgl`) 사용
- 간편한 규칙 설계환경
- 규칙 설계법 꼼꼼히 문서화
- [hangulize.org](http://hangulize.org) 개편

## 지원하는 언어

```
LANG     STAGE    ENG                      KOR
aze      draft    Azerbaijani              아제르바이잔어
bel      draft    Belarusian               벨라루스어
bul      draft    Bulgarian                불가리아어
cat      draft    Catalan                  카탈로니아어
ces      draft    Czech                    체코어
cym      draft    Welsh                    웨일스어
deu      draft    German                   독일어
ell      draft    Greek                    그리스어
epo      draft    Esperanto                에스페란토어
est      draft    Estonian                 에스토니아어
fin      draft    Finnish                  핀란드어
grc      draft    Ancient Greek            고대 그리스어
hbs      draft    Serbo-Croatian           세르보크로아트어
hun      draft    Hungarian                헝가리어
isl      draft    Icelandic                아이슬란드어
ita      draft    Italian                  이탈리아어
jpn      draft    Japanese                 일본어
kat-1    draft    Georgian (1st scheme)    조지아어(제1안)
kat-2    draft    Georgian (2nd scheme)    조지아어(제2안)
lat      draft    Latin                    라틴어
lav      draft    Latvian                  라트비아어
lit      draft    Lithuanian               리투아니아어
mkd      draft    Macedonian               마케도니아어
nld      draft    Dutch                    네덜란드어
pol      draft    Polish                   폴란드어
por      draft    Portuguese               포르투갈어
por-br   draft    Brazilian Portuguese     브라질 포르투갈어
ron      draft    Romanian                 루마니아어
rus      draft    Russian                  러시아어
slk      draft    Slovak                   슬로바키아어
slv      draft    Slovenian                슬로베니아어
spa      draft    Spanish                  스페인어
sqi      draft    Albanian                 알바니아어
swe      draft    Swedish                  스웨덴어
tur      draft    Turkish                  터키어
ukr      draft    Ukrainian                우크라이나어
vie      draft    Vietnamese               베트남어
wlm      draft    Middle Welsh             웨일스어(중세)
```

## 만든이

- 이흥섭, Heungsub Lee <<sub@subl.ee>>

## 라이선스

Hangulize 2는 BSD-3-Clause 라이선스 하에 공개되어 있습니다. 소스코드를 사용할
경우 라이선스 내용을 준수해주세요. 라이선스 전문은 `LICENSE` 파일에서 확인하실
수 있습니다.
