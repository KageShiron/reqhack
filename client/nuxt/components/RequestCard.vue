<template xmlns:v-clipboard="http://www.w3.org/1999/xhtml">
  <div class="card" >
    <header>
      <span class="http-method">{{ item.method }}</span>
      <span>{{ item.scheme }}://{{ item.host }}{{
        ((item.scheme === "http" && item.server_port === 80) || (item.scheme === "https" && item.server_port === 443)) ? "" : ":" + item.server_port
      }}{{ item.requesturi }}</span>
      <span class="from-ip-title">From:</span>
      <div class="from-ip">
        <b-tooltip :label="'client port :' + item.user_port">
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
            <div v-if="headerActiveTab === 'raw'"><b-input
              :value="item.rawrequest"
              type="textarea"
              readonly/></div>
          </div>
          <div class="column">
            <div class="info-header">
              <h3>Body( {{ item.body_length || 0 }} bytes )</h3>
              <b-field grouped>
                <b-select
                  placeholder="Viewer"
                  size="is-small">
                  <optgroup label="Text">
                    <option value="text">PlainText</option>
                    <option value="json">JSON</option>
                    <option value="form">Form</option>
                    <option value="xml">XML</option>
                  </optgroup>
                  <optgroup label="Binary">
                    <option value="image">Image</option>
                    <option value="image">Hex</option>
                  </optgroup>
                </b-select>
                <b-field>
                  <p class="control">
                    <b-dropdown>
                      <button
                        slot="trigger"
                        size="is-small"
                        class="button is-primary">
                        <span>Encode/Decode</span>
                        <b-icon icon="menu-down"/>
                      </button>
                      <b-dropdown-item @click="bodyAction('decodeURI')">decodeURI</b-dropdown-item>
                      <b-dropdown-item @click="bodyAction('decodeURIComponent')">decodeURIComponent</b-dropdown-item>
                      <b-dropdown-item @click="bodyAction('atob')">decode Base64 (atob)</b-dropdown-item>
                      <b-dropdown-item :separator="true"/>
                      <b-dropdown-item @click="bodyAction('encodeURI')">encodeURI</b-dropdown-item>
                      <b-dropdown-item @click="bodyAction('encodeURIComponent')">encodeURIComponent</b-dropdown-item>
                      <b-dropdown-item @click="bodyAction('btoa')">encode Base64 (btoa)</b-dropdown-item>
                      <b-dropdown-item :separator="true"/>
                      <b-dropdown-item @click="bodyAction('reset')">Reset</b-dropdown-item>
                    </b-dropdown>
                  </p>
                  <p class="control">
                    <a
                      v-clipboard:copy="body"
                      v-clipboard:success="onCopySuccess"
                      v-clipboard:error="onCopyError"
                      target="_blank"
                      class="button"><i 
                        class="fa fa-clipboard" 
                        custom-size="mdi-18px" /></a>
                  </p>
                  <p class="control">
                    <a
                      :href="'https://censys.io/ipv4/'+item.remoteaddr"
                      target="_blank"
                      class="button">
                      <b-icon
                        icon="cloud-download"
                        custom-size="mdi-18px" />
                    </a>
                  </p>
                </b-field>
              </b-field>
            </div>
            <div>
              <b-input
                :value="body"
                type="textarea"
                readonly/>
            </div>
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
  h3 {
    margin-right: 1rem;
    font-size: 1rem;
    color: #666;
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
</style>
<script>
export default {
  props: {
    item: {
      type: Object,
      default: null
    }
  },
  data() {
    return {
      headerActiveTab: 'table',
      bodyActionStack: [atob]
    }
  },
  computed: {
    body(vm) {
      while (true) {
        try {
          return vm.bodyActionStack.reduce((x, y) => y(x), this.item.body)
        } catch {
          vm.bodyActionStack.pop()
        }
      }
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
        message: 'Copied body data',
        type: 'is-success',
        position: 'is-top-right'
      })
    },
    onCopyError(e) {
      this.$toast.error('Copied Failed...', { duration: 3000 })
    },
    bodyAction(e) {
      switch (e) {
        case 'decodeURI':
          this.bodyActionStack.push(decodeURI)
          break
        case 'decodeURIComponent':
          this.bodyActionStack.push(decodeURIComponent)
          break
        case 'atob':
          this.bodyActionStack.push(atob)
          break
        case 'encodeURI':
          this.bodyActionStack.push(encodeURI)
          break
        case 'encodeURIComponent':
          this.bodyActionStack.push(encodeURIComponent)
          break
        case 'btoa':
          this.bodyActionStack.push(btoa)
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
