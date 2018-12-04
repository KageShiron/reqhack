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
            <span class="method">{{ item.method }}</span>
            <span>{{ item.host }}{{ item.requesturi }}</span>
            <time :datetime="new Date(item.time).toISOString()">{{ item.time }}</time>
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
  }
  .content {
    padding: 0.5em;
  }
}

.http-method {
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 3px 5px;
  margin-right: 1em;
}
h1 {
  padding: 1em;
}
</style>
