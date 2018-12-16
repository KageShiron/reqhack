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
              <span>{{ item.scheme }}://{{ item.host }}{{
                ((item.scheme === "http" && item.port === 80) || (item.scheme === "https" && item.port === 443)) ? "" : ":" + item.port
              }}{{ item.requesturi }}</span>
              <span class="from-ip-title">From:</span>
              <div class="from-ip">
                <span>
                  {{ item.remoteaddr }}
                </span>
                <div class="ip-actions">
                  <div class="field has-addons">
                    <p class="control">
                      <b-tooltip label="Copy remote IP address">
                        <a
                          v-clipboard:copy="item.remoteaddr"
                          v-clipboard:success="onCopySuccess"
                          v-clipboard:error="onCopyError"
                          target="_blank"
                          class="button"><i class="fa fa-clipboard" /></a>
                      </b-tooltip>
                    </p>
                    <p class="control">
                      <b-tooltip label="View in Censys">
                        <a
                          :href="'https://censys.io/ipv4/'+item.remoteaddr"
                          target="_blank"
                          class="button"><img src="https://censys.io/favicon.ico"></a>
                      </b-tooltip>
                    </p>
                    <p class="control">
                      <b-tooltip label="View in Shodan">
                        <a
                          :href="'https://www.shodan.io/search?query='+item.remoteaddr"
                          class="button"><img src="https://static.shodan.io/shodan/img/favicon.png"></a>
                      </b-tooltip>
                  </p></div>
                </div>

              </div>
              <time :datetime="item.time">{{ $moment(item.time).format("YYYY/MM/DD HH:mm:ss Z") }} <br> {{ $moment(item.time).fromNow() }}</time>
            </header>
            <div class="content">
              <div class="container">
                <div class="columns">
                  <div class="column">
                    <div class="httpheaders">
                      <h3>HTTP Headers</h3>
                      <b-field>
                        <b-radio-button
                          v-model="headerActiveTab"
                          size="is-small"
                          native-value="table">
                          <b-icon 
                            icon="table" 
                            custom-size="mdi-18px"/>Table
                        </b-radio-button>
                        <b-radio-button
                          v-model="headerActiveTab"
                          size="is-small"
                          native-value="json">
                          <b-icon 
                            icon="json"
                            custom-size="mdi-18px"/>JSON
                        </b-radio-button>
                        <b-radio-button
                          v-model="headerActiveTab"
                          size="is-small"
                          native-value="raw">
                          <b-icon 
                            icon="format-columns"
                            custom-size="mdi-18px"/>Raw
                        </b-radio-button>
                      </b-field>
                    </div>
                    <div v-if="headerActiveTab === 'table'">
                      <table>
                        <tr
                          v-for="(v,k) in item.header"
                          :key="k">
                          <th>{{ k }}</th>
                          <td>{{ v.join("\n") }}</td>
                        </tr>
                      </table>
                    </div>
                    <div v-if="headerActiveTab === 'json'"><b-input
                      :value="JSON.stringify(item.header,null,2)"
                      type="textarea"
                      readonly/></div>
                    <div v-if="headerActiveTab === 'raw'">{{ item.rawrequest }}</div>
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
    headerActiveTab: {
      get() {
        return this.$store.state.bin.headerActiveTab
      },
      set(v) {
        this.$store.commit('bin/update_headerActiveTab', v)
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

nav.tabs {
  & > ul {
    margin: 0;
  }

  &::before {
    content: 'HTTP Headers';
    font-size: 1rem;
    font-weight: bold;
    color: #666;
  }
}

.httpheaders {
  display: flex;
  h3 {
    margin-right: 1rem;
  }
}
</style>
