<template xmlns:v-clipboard="http://www.w3.org/1999/xhtml">
  <div class="card" >
    <header>
      <span class="http-method">{{ item.method }}</span>
      <b-field class="url">
        <b-input
          :value="url"
          readonly
          class="url-text"
          type="text" />
        <p
          class="control">
          <a
            v-clipboard:copy="url"
            v-clipboard:success="onCopyBodySuccess"
            v-clipboard:error="onCopyError"
            target="_blank"
            class="button">
          <i class="fa fa-clipboard mdi-18px" /></a>
        </p>
      </b-field>
      <span class="from-ip-title">From:</span>
      <div class="from-ip">
        <b-tooltip :label="'Client port :' + item.user_port">
          {{ item.remoteaddr }}
        </b-tooltip>
        <div class="ip-actions">
          <div class="field has-addons">
            <p class="control">
              <b-tooltip label="Copy remote IP address">
                <a
                  v-clipboard:copy="item.remoteaddr"
                  v-clipboard:success="onCopyBodySuccess"
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
                  class="button"><img src="/icon/censys.ico"></a>
              </b-tooltip>
            </p>
            <p class="control">
              <b-tooltip label="View in Shodan">
                <a
                  :href="'https://www.shodan.io/search?query='+item.remoteaddr"
                  target="_blank"
                  class="button"><img src="/icon/shodan.png"></a>
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
            <div class="info-header">
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
              <table class="http-header-table">
                <tr
                  v-for="(v,k) in item.header"
                  :key="k">
                  <td><a
                    :href="'https://developer.mozilla.org/docs/Web/HTTP/Headers/' + k"
                    title="View in MDN"
                    target="_blank"><img src="/icon/mdn.svg"></a></td>
                  <th>{{ k }}</th>
                  <td>{{ v.join("\n") }}</td>
                </tr>
              </table>
            </div>
            <div v-if="headerActiveTab === 'json'"><b-input
              :value="JSON.stringify(item.header,null,2)"
              type="textarea"
              readonly/></div>
            <div v-if="headerActiveTab === 'raw'"><b-input
              :value="item.rawrequest"
              type="textarea"
              readonly/></div>
            <div class="snippet">
              <div class="info-header">
                <h3>Snippet(β)</h3>
                <b-field grouped>
                  <b-select
                    v-model="snippetSelected"
                    size="is-small">
                    <option
                      v-for="rule in snippetsRules"
                      :value="rule"
                      :key="rule">
                      {{ rule }}
                    </option>
                  </b-select>
                  <b-field>
                    <p class="control">
                      <a
                        v-clipboard:copy="snippet"
                        v-clipboard:success="onCopyBodySuccess"
                        v-clipboard:error="onCopyError"
                        :disabled="disabled"
                        target="_blank"
                        class="button">
                        <b-icon
                          icon="clipboard"
                          pack="fa" />
                      </a>
                    </p>
                    <p class="control">
                      <a
                        :disabled="disabled"
                        class="button"
                        @click="downloadSnippet">
                        <b-icon
                          icon="cloud-download"
                          custom-size="mdi-18px" />
                      </a>
                    </p>
                  </b-field>
                </b-field>
              </div>
              <b-input
                :value="snippet"
                type="textarea"
                readonly/>
            </div>
          </div>
          <div class="column">
            <BodyView
              :size="item.body_length"
              :body="item.body"
              :mime="item.header['Content-Type'].join('')"
              name="Body" />
            <!-- todo: -->
            <div class="info-header">
              <h3>Form</h3>
            </div>
            <table>
              <tr
                v-for="key in Object.keys(item.form)"
                :key="key">
                <th>{{ key }}</th>
                <td><b-input
                  :value="item.form[key]"
                  readonly /></td>
              </tr>
            </table>
            <BodyView
              v-for="mfile in item.multipartform"
              :key="mfile.Filename"
              :name="mfile.Filename"
              :size="mfile.Size"
              :body="mfile.Body"
              :mime="mfile.Header['Content-Type'].join('')"
              icon="file"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style  lang="scss">
.textarea {
  height: 15em !important;
}
a[disabled] {
  pointer-events: none;
}
.card {
  margin-bottom: 1em;

  header {
    padding: 0.5em;
    border-bottom: 1px solid #ccc;
    display: grid;
    grid-template-columns: auto 1fr auto;
    grid-gap: 5px;
    grid-template-rows: auto auto;

    & > :last-child {
      margin-left: auto;
    }

    .http-method {
      background-color: #666;
      color: white;
      border-radius: 3px;
      padding: 2px 5px;
      max-width: 100px;
      word-break: break-all;
      height: 24px;
      align-self: center;
      font-size: 13px;
    }

    .url {
      margin: 0;
      & > * {
        height: 1.8rem;
      }
      .url-text {
        flex-grow: 1;
        input {
          padding: 0.1rem 0.3rem;
          height: 1.8rem;
        }
      }

      p a {
        padding: 0.1rem 0.5rem;
        height: 1.8rem;
      }
    }

    .from-ip-title {
      align-self: center;
      font-size: 13px;
    }

    time {
      grid-column: 3;
      grid-row: 1/3;
      text-align: right;
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
  h3 {
    margin-right: 0.3rem;
    font-size: 0.9rem;
    color: #666;
    border-top: 2px solid #714dd2;
    padding: 3px 10px;
    margin-bottom: 0;
    box-shadow: 0 0 3px #ccc;
    height: 27px;
  }
  .info-header {
    margin-bottom: 0.3rem;
    display: flex;
    button,
    a {
      font-size: 0.75rem;
    }
    .control {
      margin: 0;
    }
  }
}

table.http-header-table {
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

    & > td:first-child {
      min-width: 20px;
      padding: 0 !important;
      a {
        vertical-align: middle !important;
        img {
          opacity: 0.5;
          height: 16px;
          width: 16px;
          display: block;
          &:hover {
            opacity: 0.8;
          }
        }
      }
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
.snippet {
  margin-top: 1em;
}
</style>
<script>
import BodyView from './BodyView'

export default {
  components: { BodyView },
  props: {
    item: {
      type: Object,
      default: null
    }
  },
  data() {
    return {
      headerActiveTab: 'table',
      bodyActionStack: [atob],
      viewer: 'text',
      snippetSelected: 'cURL'
    }
  },
  computed: {
    body() {
      let res = this.bodyActionStack.reduce((x, y) => y(x), this.item.body)
      if (this.viewer === 'json') {
        try {
          return JSON.stringify(JSON.parse(res), null, 2)
        } catch (e) {
          return 'JSON Parse Error.\n' + e
        }
      } else if (this.viewer === 'xml') {
        const dom = new DOMParser().parseFromString(res, 'application/xml')
        if (dom.documentElement.nodeName === 'parsererror') {
          return dom.documentElement.childNodes[0].nodeValue
        }
        return new XMLSerializer().serializeToString(dom)
      }

      return res
    },

    url() {
      return `${this.item.scheme}://${this.item.host}${
        (this.item.scheme === 'http' && this.item.server_port === 80) ||
        (this.item.scheme === 'https' && this.item.server_port === 443)
          ? ''
          : ':' + this.item.server_port
      }${this.item.requesturi}`
    },
    form() {
      void this.headerActiveTab // なぜか消すと動かない
      try {
        return new URLSearchParams(this.body).entries()
      } catch (e) {
        return {}
      }
    },
    disabled() {
      console.log(this.item.body_length)
      return this.item.body_length === 0
    },
    snippet() {
      switch (this.snippetSelected) {
        case 'cURL':
          let headers = ''
          for (const [k, vs] of Object.entries(this.item.header)) {
            for (const v of vs) {
              headers += `-H "${k}:${v}" `
            }
          }
          return `curl -X ${this.item.method} ${headers} --data ${atob(
            this.item.body
          )} ${this.url}`
        default:
          return ''
      }
    },
    snippetsRules() {
      return ['cURL']
    }
  },
  methods: {
    onCopySuccess(e) {
      this.$toast.open({
        message: 'Copied ' + e.text,
        type: 'is-success',
        position: 'is-top-right'
      })
    },
    onCopyBodySuccess(e) {
      this.$toast.open({
        message: 'Copied data.',
        type: 'is-success',
        position: 'is-top-right'
      })
    },
    onCopyError(e) {
      this.$toast.error('Copied Failed...', { duration: 3000 })
    },
    download() {
      var blob = new Blob([this.body], { type: 'text/plain' })
      let link = document.createElement('a')
      link.href = window.URL.createObjectURL(blob)
      link.download = 'body.txt'
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
    },
    downloadSnippet() {
      var blob = new Blob([this.body], { type: 'text/plain' })
      let link = document.createElement('a')
      link.href = window.URL.createObjectURL(blob)
      link.download = 'snippet.txt'
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
    },
    addAction(newValue) {
      try {
        newValue(this.bodyActionStack.reduce((x, y) => y(x), this.item.body))
        this.bodyActionStack.push(newValue)
      } catch {}
    },
    bodyAction(e) {
      switch (e) {
        case 'decodeURI':
          this.addAction(decodeURI)
          break
        case 'decodeURIComponent':
          this.addAction(decodeURIComponent)
          break
        case 'atob':
          this.addAction(atob)
          break
        case 'encodeURI':
          this.addAction(encodeURI)
          break
        case 'encodeURIComponent':
          this.addAction(encodeURIComponent)
          break
        case 'btoa':
          this.addAction(btoa)
          break
        case 'reset':
          this.bodyActionStack = [atob]
          break
      }
      console.log(this.bodyActionStack)
    }
  }
}
</script>
