<template>
  <section class="section">
    <div class="field is-horizontal binurl">
      <span>{{ protocol }}//</span>
      <div class="control has-icons-right">
        <input 
          v-model="binname"
          class="input"
          type="text"
        >
      </div>
      <span>.{{ host }}</span>

    </div>
    <div class="field">
      <button 
        class="button is-primary is-right" 
        @click="$store.dispatch('createBin', binname)">CreateBin</button>
    </div>

    <div v-if="created">
      {{ error }}
    </div>
    <div class="news">
      不具合、ご意見、ご感想などはSNSなどのDirect Messageでお願いします。
      プレゼンや卒論用にユースケースも集めたいので、こんな使い方しましたみたいなものもお待ちしております！
      <div class="container">
        <h2 class="title">使い方</h2>
        <ol>
          <li>好きな英数字を入れて、CreateBinボタン</li>
          <li>解析ページにジャンプするので、Logger URLを別のタブやcurlなどでアクセス</li>
          <li>解析ページを更新するとリクエスト内容を確認できます</li>
        </ol>
      </div>
      <div class="container">
        <h2 class="title">注意点</h2>
        <ul>
          <li>現在はbinの閲覧制限が無いので、機密情報を残さないことをおすすめします。削除対応は@KageShironまで直接ご連絡を。</li>
          <li>データは2日ぐらいで消えるはずですが、うまく消えなかったり、事故って突然DBが消しとぶ場合があります。</li>
          <li>開発用サーバーなので、予告なく止まったりぶっ壊れたりします</li>
        </ul>
      </div>
      <div class="container">
        <h2 class="title">最近の更新</h2>
        <ul>
          <li>
            2018/12/24
            <ul>
              <li>サーバーポートが間違っている不具合を修正</li>
              <li>クライアントIPアドレスのToolTipでクライアントポートを確認できるように</li>
            </ul>
          </li>
          <li>内部公開しました</li>
        </ul>
      </div>
      <div class="container">
        <h2 class="title">既知の不具合</h2>
        <ul>
          <li>一時的にhttpsが使えません</li>
        </ul>
      </div>
      <div class="container">
        <h2 class="title">近日対応予定</h2>
        <ul>
          <li>既知の不具合の修正</li>
          <li>curlでコピー</li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script>
import Logo from '~/components/Logo.vue'

export default {
  components: {
    Logo
  },
  computed: {
    binname: {
      get() {
        return this.$store.state.binname
      },
      set(value) {
        this.$store.commit('updateBinName', value)
      }
    },
    host() {
      return location.host
    },
    protocol() {
      return location.protocol
    },
    created() {
      return this.$store.state.created
    }
  }
}
</script>

<style>
section {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.binurl {
  display: flex;
  align-items: center;
}
.news h2.title {
  margin: 1em 0 0.2em 0 !important;
}
.news ul {
  list-style: disc;
}
</style>
