<template>
  <div>
    <h1 class="title">Inspect {{ $route.params.name }}</h1>
    <a :href="`${protocol}//${$route.params.name}.${host}/`">{{ protocol }}//{{ $route.params.name }}.{{ host }}</a>


    <div class="columns">
      <div class="column">
        <div
          v-for="(item,key) in items"
          :key="key"
          class="card"
        >
          <header>
            <span class="http-method">{{ item.method }}</span>
            <span>{{ item.host }}{{ item.requesturi }}</span>
            <span class="from-ip-title">From:</span>
            <div class="from-ip">{{ item.remoteaddr }}
              <a
                :href="'https://censys.io/ipv4/'+item.remoteaddr"
                class="button"><img src="https://censys.io/favicon.ico"></a>
              <a
                :href="'https://www.shodan.io/search?query='+item.remoteaddr"
                class="button"><img src="https://static.shodan.io/shodan/img/favicon.png"></a></div>
            <time :datetime="item.time">{{ $moment(item.time).format("YYYY/MM/DD HH:mm:ss Z") }} <br> {{ $moment(item.time).fromNow() }}</time>
          </header>
          <div class="content">
            <div class="field has-addons">
              <p class="control">
                <a class="button">
                  <span>Table</span>
                </a>
              </p>
              <p class="control">
                <a class="button">
                  <span>Raw</span>
                </a>
              </p>
              <p class="control">
                <a class="button">
                  <span>Json</span>
                </a>
              </p>
            </div>
            <table>
              <h2>Http Headers</h2>
              <tr
                v-for="(v,k) in item.header"
                :key="k"><th>{{ k }}</th><td>{{ v.join("\n") }}</td></tr>
            </table>
            <div>
              <input type="text" >
              <button/>
            </div>
          </div>
          <!--
          <header class="card-header">
            <p class="card-header-title">
              <span class="http-method">{{ item.method }}</span>
              <span>{{ item.requesturi }}</span>
            </p>
            <p class="card-header-title">
              <span>{{ item.remoteaddr }}</span>
            </p>
          </header>
          <div class="card-content">
            <div class="content">
              {{ new Date(item.time).toISOString() }}
              <table>
                <tr
                  v-for="(v,k) in item.header"
                  :key="k"><th>{{ k }}</th><td>{{ v }}</td></tr>
              </table>
            </div>
          </div>-->
    </div></div></div>
  </div>
</template>

<script>
import moment from 'moment'
export default {
  async fetch({ store, params }) {
    await store.dispatch('bin/fetch_bin', params.name)
  },
  computed: {
    items() {
      return this.$store.state.bin.items[this.$route.params.name]
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
<style lang="scss">
.card {
  margin: 1em;
  header {
    padding: 0.5em;
    border-bottom: 1px solid #ccc;
    display: grid;
    grid-template-columns: min-content max-content auto;
    grid-template-rows: 1fr 1fr;
    & > :last-child {
      margin-left: auto;
    }
    .http-method {
      background-color: #666;
      color: white;
      border-radius: 3px;
      padding: 2px 5px;
      margin-right: 0.5em;
    }
    time {
      grid-column: 3;
      grid-row: 1/2;
    }
    .from-ip .button {
      padding: 0.3em;
      line-height: 1em;
      height: auto;
      img {
        height: 1em;
      }
    }
  }
  .content {
    padding: 0.5em;
  }
}

h1 {
  padding: 1em;
}
</style>
