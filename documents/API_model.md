# 絵文字のパラメータついて
主に絵文字について

## FaceAnotationの量について
https://cloud.google.com/vision/docs/reference/rest/v1/AnnotateImageResponse?hl=ja#Likelihood
ここにアノテーションとかの型が載っていた


### 感情のタイプ
- joyLikelihood	
    - Joy likelihood.
    - 喜び
- sorrowLikelihood	
    - Sorrow likelihood.
    - 悲しみ
- angerLikelihood	
    - Anger likelihood.
    - 怒り
- surpriseLikelihood	
    - Surprise likelihood.
    - 驚き
- underExposedLikelihood	
    - Under-exposed likelihood.
    - 露光不足みたいな
    - 基本的には使えなさそう
- blurredLikelihood	
    - Blurred likelihood.
    - 霞んでるかどうか？
    - これも使えなさそうな気がするが...
- headwearLikelihood	
    - Headwear likelihood.
    - 帽子かぶってても表情に関係ないので，このパラメータが100でそれ以外がゼロの時は👒を表示させるとかでもいいかも

### Likelihoodのタイプ
| Type           | 説明 | 絵文字としての値 |
| --------       | --------               | ----- |
| UNKNOWN        | Unknown likelihood.    | 0     |
| VERY_UNLIKELY  | It is very unlikely.   | 0     |
| UNLIKELY       | It is unlikely.        | 20    |
| POSSIBLE       | It is possible.        | 50    |
| LIKELY         | It is likely.          | 70    |
| VERY_LIKELY    | It is very likely.     | 100   |

## コメント
基本的には4つのパラメータ(喜び，怒り，悲しみ，驚き)で決まりそう．少なくても絵文字は4~10個くらいになるのかな．

## まいまいがデザインした絵文字と対応(適当)
(喜び，怒り，悲しみ，驚き)
- 楽しい -> (100, 0, 0, 0)
- ニコニコ -> (50-70, 0, 0, 0)
- Love -> (100, 0, 0, 0)(楽しいとランダムで表示させてもいいかも．あと色とか)
- ぴえん -> (0-50, 0, 50-100, 0)
- 驚き -> (0, 0, 0, 50)
- 驚き，かみなり -> (0, 0, 0, 100)
- 😭 -> (0, 0, 100, 0)
- 😡 -> (0, 100, 0, 0)
- 😠 -> (0, 50-70, 0, 0)

![](https://i.imgur.com/wUcTgif.jpg)
