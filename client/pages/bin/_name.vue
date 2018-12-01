<template>
  <div>
    <h1>Inspect {{ $route.params.name }}</h1>


    <div class="columns">
      <div class="column">
        <div
          v-for="(item,key) in items"
          :key="key"
          class="card"
        >
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
          </div>
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
    }
  }
}
</script>
<style>
.card {
  margin: 1em;
}

.http-method {
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 3px 5px;
}
h1 {
  padding: 1em;
}
</style>
