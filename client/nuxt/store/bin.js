import axios from 'axios'

export const state = () => ({
  items: {},
  headerActiveTab: 'table'
})

export const mutations = {
  update_bin(state, value) {
    if (!value.name) throw 'Bad update bin'
    state.items[value.name] = Object.assign(
      {},
      state.items[value.name],
      value.data
    )
    if (value.data.secret) {
      let secrets = JSON.parse(localStorage.getItem('secrets') || '{}')
      secrets[value.name] = value.data.secret
      localStorage.setItem('secrets', JSON.stringify(secrets))
    }
  },
  update_headerActiveTab(state, value) {
    state.headerActiveTab = value
  }
}

export const actions = {
  async fetch_bin(context, params) {
    const lsSecret = JSON.parse(localStorage.getItem('secrets') || '{}')[
      params.name
    ]
    const secret =
      params.secret ||
      (context.state.items[params.name] &&
        context.state.items[params.name].secret) ||
      lsSecret ||
      ''
    const { data } = await axios.get(
      `/v1/${params.name}/items?secret=${secret}`
    )
    console.log(data)
    context.commit('update_bin', {
      name: params.name,
      data: { data: data.reverse() }
    })
  }
}
