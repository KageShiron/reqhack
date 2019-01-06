<template xmlns:v-clipboard="http://www.w3.org/1999/xhtml">
  <div>
    <section class="hero">
      <div class="hero-body">
        <div class="container">
          <h1 class="title"><span style="font-size:80%;color:#666;">Inspect</span> {{ $route.params.name }}</h1>
          <h2 class="subtitle">
            Logger URL: <a :href="`${protocol}//${$route.params.name}.${host}/`">{{ protocol }}//{{ $route.params.name }}.{{ host }}</a>

            <a
              v-clipboard:copy="`${protocol}//${$route.params.name}.${host}/`"
              v-clipboard:success="onCopySuccess"
              v-clipboard:error="onCopyError"
              target="_blank"
              class="button">
              <i class="fa fa-clipboard" />
            </a>
          </h2>
        </div>
      </div>
    </section>

    <div class="container">
      <div class="columns">
        <div class="column">
          <RequestCard 
            v-for="item in items"
            :key="item.id"
            :item="item"/>
        </div>
      </div>
    </div>
</div></template>

<script>
import moment from 'moment'
import RequestCard from '../../components/RequestCard'
export default {
  components: { RequestCard },
  async fetch({ store, params,query }) {
    await store.dispatch('bin/fetch_bin', {
      name: params.name,
      secret: query.secret
    })
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
  head() {
    return {
      title: this.$route.params.name + ' - reqhack'
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
    onCopyError(e) {
      this.$toast.error('Copied Failed...', { duration: 3000 })
    }
  }
}
</script>
<style lang="scss">
.hero-body {
  padding-top: 1em;
  padding-bottom: 1em;
}
.hero-subtitle {
  font-size: 1rem;
}
.button {
  vertical-align: baseline;
}
</style>
