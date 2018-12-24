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
</style>
