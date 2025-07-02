import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import styles from "./page.module.css";

const Code = ({ children }: { children: string }) => (
  <SyntaxHighlighter language="python" className={styles.code}>
    {children}
  </SyntaxHighlighter>
);

const DocumentsPage: React.FC = async () => {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>🐍 Python 文法まとめ</h1>

      <section className={styles.section}>
        <h2 className={styles.heading}>1. 変数とデータの種類（データ型）</h2>
        <p className={styles.text}>
          数字や文字などを入れる「箱」を変数と呼びます。
        </p>
        <Code>{`x = 10        # 整数（int）
y = 3.5       # 小数（float）
name = "Tom"  # 文字列（str）
is_ok = True  # 真偽値（bool）`}</Code>
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>2. よく使う演算子</h2>
        <p className={styles.text}>計算や比較に使います。</p>
        <Code>{`# 算術演算
a + b     # 足し算
a - b     # 引き算
a * b     # 掛け算
a / b     # 割り算（小数）
a // b    # 割り算（整数）
a % b     # 剰余（あまり）

# 比較（if文などでよく使う）
a == b    # a と b が等しい
a != b    # a と b が異なる
a < b     # a は b より小さい
a <= b    # a は b 以下
a > b     # a は b より大きい
a >= b    # a は b 以上

# 論理演算
and, or, not`}</Code>
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>3. 条件分岐（if文）</h2>
        <p className={styles.text}>
          「もし〜ならこうする」という処理をしたいときに使います。<br />
          たとえば、「もし年齢が18歳以上なら成人です」と判断するときの例です。
        </p>
        <Code>{`age = 20

if age >= 18:
    # age が 18 以上ならここが実行される（成人）
    adult = True
else:
    # age が 18 未満ならここが実行される（未成年）
    adult = False

# adult は True`}</Code>
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>4. 繰り返し（ループ）</h2>
        <p className={styles.text}>
          同じ処理を何度も繰り返したいときに使います。<br />
          「for」はリストの中身を1つずつ取り出して繰り返します。<br />
          「while」は条件が真の間、繰り返します。<br />
          <strong>break</strong> はループを途中で終わらせたいときに使います。
        </p>
        <Code>{`numbers = [10, 20, 30]

# numbers の中の値を1つずつ n に入れて処理する
for n in numbers:
    doubled = n * 2
    # doubled は 20, 40, 60 になる

# よく使う繰り返しの書き方
# 0から4の中の値を1つずつ i に入れて処理する
for i in range(5):
    # 0, 1, 2, 3, 4 と順に処理される
    if i == 3:
        break  # i が 3 のときにループを途中で抜ける

# while は「条件を満たす間」繰り返すときに使う
count = 0
while count < 3:
    count += 1
    # count は 1, 2, 3 と増える`}</Code>
      </section>


      <section className={styles.section}>
        <h2 className={styles.heading}>5. 関数</h2>
        <p className={styles.text}>
          よく使う処理をまとめて名前をつけたものが関数です。<br />
          呼ぶだけで何度でも使えます。
        </p>
        <Code>{`def greet(name):
    # name に渡された名前を使って挨拶を作る
    message = "こんにちは、" + name + "さん"
    return message

result = greet("太郎")
# result は "こんにちは、太郎さん"`}</Code>
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>6. リスト（配列）</h2>
        <p className={styles.text}>
          複数の値をまとめて扱う箱です。順番があり、0番目から数えます。
        </p>
        <Code>{`fruits = ["りんご", "バナナ", "みかん"]

fruits.append("ぶどう")
# fruits は ["りんご", "バナナ", "みかん", "ぶどう"] になる

first = fruits[0]
# first は "りんご"

for fruit in fruits:
    # fruits の中身を1つずつ fruit に入れて処理する`}</Code>
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>7. 辞書（連想配列）</h2>
        <p className={styles.text}>
          キー（名前）と値をセットで扱える箱です。
        </p>
        <Code>{`person = {"名前": "アミ", "年齢": 20}

name = person["名前"]
# name は "アミ"

person["年齢"] = 21
# 年齢が 21 に更新される

for key, value in person.items():
    # キーと値を順に取り出して処理できる`}</Code>
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>8. よく使う組み込み関数</h2>
        <p className={styles.text}>
          Python に最初から用意されている便利な関数です。
        </p>
        <Code>{`range(n)        # 0 から n-1 までの連続した整数を作る
len(list)       # リストや文字列の要素数を調べる
sum(list)       # リストの中の数の合計を計算する
enumerate(list) # 順番（インデックス）付きでリストを繰り返す`}</Code>
      </section>

      <section className={styles.section}>
        <h2 className={styles.heading}>🎯 総まとめ</h2>
        <p className={styles.text}>
          🧠 <strong>すべての文法を覚える必要はありません</strong>。<br />
          （ここに書いてるのはいつかは覚えて欲しいけど）<br />
          分からないことがあったら、<strong>「python 調べたいこと」</strong>で検索しましょう。<br />
          コーディングは「覚える」より「調べて試す」ことが大切です。
        </p>
      </section>
    </div>
  );
};

export default DocumentsPage;
