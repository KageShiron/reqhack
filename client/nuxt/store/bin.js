import axios from 'axios'

export const state = () => ({
  items: {},
  headerActiveTab: 'table'
})

export const mutations = {
  update_bin(state, value) {
    if(!value.name) throw "Bad update bin";
    state.items[value.name] = Object.assign(value.data,state.items[value.name])
  },
  update_headerActiveTab(state, value) {
    state.headerActiveTab = value
  }
}

export const actions = {
  async fetch_bin(context, params) {
<<<<<<< HEAD
    const secret = params.secret || ( context.state.items[params.name] && context.state.items[params.name].secret)
    const { data } = await axios.get(
      `/v1/${params.name}/items?secret=${secret}`
    )
    context.commit('update_bin', {
      name: params.name,
      data: { data: data.reverse() }
    })
=======
    const { data } = await axios.get(
      `/v1/${params.name}/items?secret=${params.secret}`
    )
    context.commit('update_bin', { nama: params.name, data: data.reverse() })
>>>>>>> 94bdb1605975256d4634df4e4eb068db4ec9fc83
  }
}
