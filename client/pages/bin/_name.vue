<template xmlns:v-clipboard="http://www.w3.org/1999/xhtml">
  <div>
    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">Inspect {{ $route.params.name }}</h1>
          <h2 class="subtitle">
            Logger URL: <a :href="`${protocol}//${$route.params.name}.${host}/`">{{ protocol }}//{{ $route.params.name }}.{{ host }}</a>
          </h2>
        </div>
      </div>
    </section>

    <div class="container">
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
              <div class="from-ip">
                <span>
                  {{ item.remoteaddr }}
                </span>
                <div class="ip-actions">
                  <div class="field has-addons">
                    <p class="control">
                      <a
                        v-clipboard:copy="item.remoteaddr"
                        v-clipboard:success="onCopySuccess"
                        v-clipboard:error="onCopyError"
                        target="_blank"
                        data-tooltip="Copy remote IP address"
                        class="button tooltip"><i class="fa fa-clipboard" /></a>
                    </p>
                    <p class="control">
                      <a
                        :href="'https://censys.io/ipv4/'+item.remoteaddr"
                        target="_blank"
                        data-tooltip="View Censys"
                        class="button tooltip"><img src="https://censys.io/favicon.ico"></a>
                    </p>
                    <p class="control">
                      <a
                        :href="'https://www.shodan.io/search?query='+item.remoteaddr"
                        data-tooltip="View Shodan"
                        class="button tooltip"><img src="https://static.shodan.io/shodan/img/favicon.png"></a>
                  </p></div>
                </div>

              </div>
              <time :datetime="item.time">{{ $moment(item.time).format("YYYY/MM/DD HH:mm:ss Z") }} <br> {{ $moment(item.time).fromNow() }}</time>
            </header>
            <div class="content">
              <div class="container">
                <div class="columns">
                  <div class="column">
                    <h3>HTTP Headers</h3>
                    <table>
                      <tr
                        v-for="(v,k) in item.header"
                        :key="k">
                        <th>{{ k }}</th>
                        <td>{{ v.join("\n") }}</td>
                      </tr>
                    </table>
                  </div>
                  <div class="column">
                    <h3>Body</h3>
                    <div>
                      {{ item.body }}
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
</div></template>

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
  },
  methods: {
    onCopySuccess(e) {
      this.$toast.success('Copied ' + e.text, {
        duration: 3000,
        iconPack: 'fontawesome',
        icon: 'clipboard'
      })
    },
    onCopyError(e) {
      this.$toast.error('Copied Failed...', { duration: 3000 })
    }
  }
}
</script>
<style lang="scss">
.card {
  margin-bottom: 1em;
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
      grid-row: 1/3;
    }
    .from-ip .button {
      padding: 0.3em;
      line-height: 1em;
      height: auto;
      img {
        height: 1em;
      }
    }
    .ip-actions {
      display: inline-block;
    }
  }
  .content {
    padding: 0.5em;
  }
}
.hero-body {
  padding-top: 1em;
  padding-bottom: 1em;
}
.hero-subtitle {
  font-size: 1rem;
}

.content h3 {
  margin: 0;
  font-size: 1rem;
  color: #666;
  .field {
    display: inline-flex !important;
    vertical-align: text-top;
    font-size: 1rem;

    .control {
      margin-bottom: 0 !important;
      .button {
        padding: 0.2rem 0.3rem;
        height: auto;
        font-weight: normal;
      }
    }
  }
}
table {
  font-size: 0.9rem;
  tr {
    &:hover {
      background-color: #f3f3f9;
      transition: background-color 0.25s;
    }
    & > * {
      padding: 0.1rem 0.3rem !important;
      vertical-align: middle !important;
    }
  }
}
</style>
