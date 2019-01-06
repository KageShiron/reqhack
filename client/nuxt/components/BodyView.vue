<template>
  <div class="bodyview">
    <div class="info-header">
      <h3>{{ name }}</h3>
      <small>{{ size }} bytes</small>
      <b-field grouped>
        <b-select
          :disabled="disabled"
          v-model="viewer"
          placeholder="Viewer"
          size="is-small">
          <optgroup label="Text">
            <option value="text">PlainText</option>
            <option value="json">JSON</option>
            <option value="form">Form</option>
            <option value="xml">XML</option>
          </optgroup>
          <optgroup label="Binary">
            <option value="image">Image(β)</option>
            <option
              value="image"
              disabled>Hex (WIP)</option>
          </optgroup>
        </b-select>
        <b-field>
          <p class="control">
            <b-dropdown :disabled="disabled">
              <button
                slot="trigger"
                size="is-small"
                title="Encode/Decode"
                class="button">
                <b-icon
                  icon="percent"
                  pack="fa"
                  custom-size="mdi-18px" />
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
              v-clipboard:success="onCopyBodySuccess"
              v-clipboard:error="onCopyError"
              :disabled="disabled"
              target="_blank"
              class="button"
              title="Copy"
            >
              <b-icon
                icon="clipboard"
                pack="fa" />
            </a>
          </p>
          <p class="control">
            <a
              :disabled="disabled"
              class="button"
              title="Download as file"
              @click="download">
              <b-icon
                icon="cloud-download"
                custom-size="mdi-18px" />
            </a>
          </p>
        </b-field>
      </b-field>
    </div>
    <div v-if="filename !== ''">
      <i :class="'mdi mdi-' + icon"/>
      {{ filename }}
    </div>

    <div v-if="viewer !== 'form' && viewer !== 'image'">
      <b-input
        :value="changedBody"
        :disabled="disabled"
        type="textarea"
        readonly/>
    </div>
    <div v-if="viewer === 'form'">
      <table>
        <tr
          v-for="[key,val] in Array.from(form)"
          :key="key">
          <th>{{ key }}</th>
          <td><b-input
            :value="val"
            readonly /></td>
        </tr>
      </table>
    </div>
    <div v-if="viewer === 'image'">
      <img :src="image" >
    </div>
  </div>
</template>
<script>
function getDefaultViewer(mime) {
  if (mime === 'application/x-www-form-urlencoded') {
    return 'form'
  }
  if (mime === 'application/json') {
    return 'json'
  }
  if (mime.startsWith('image/')) {
    return 'image'
  }
  if (mime === 'aplication/xml') {
    return 'xml'
  }
  return 'text'
}
export default {
  props: {
    body: {
      type: String,
      default: null
    },
    size: {
      type: Number,
      default: 0
    },
    mime: {
      type: String,
      default: null
    },
    name: {
      type: String,
      default: null
    },
    filename: {
      type: String,
      default: null
    },
    icon: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      bodyActionStack: [atob],
      viewer: 'text'
    }
  },
  computed: {
    changedBody() {
      console.log(this.body)
      let res = this.bodyActionStack.reduce((x, y) => y(x), this.body)
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
    form() {
      void this.headerActiveTab // なぜか消すと動かない
      try {
        return new URLSearchParams(this.changedBody).entries()
      } catch (e) {
        return {}
      }
    },
    disabled() {
      console.log(this.size)
      return this.size === 0
    },
    image() {
      if (this.mime.startsWith('image/')) {
        return `data:${this.mime};base64,${this.body}`
      } else {
        return `data:;base64,${this.body}`
      }
    }
  },
  mounted() {
    this.viewer = getDefaultViewer(this.mime) || ''
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
      let blob = new Blob([this.body], { type: this.mime })
      let link = document.createElement('a')
      if (this.viewer === 'image') {
        link.href = this.image
      } else {
        let blob = new Blob([this.body], { type: this.mime })
        link.href = window.URL.createObjectURL(blob)
      }
      link.download = this.filename || this.name
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
    },
    addAction(newValue) {
      try {
        newValue(this.bodyActionStack.reduce((x, y) => y(x), this.changedBody))
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
<style lang="scss" scoped>
small {
  align-self: center;
  margin-right: 0.5rem;
}
.bodyview {
  margin-bottom: 1rem;
}
</style>
