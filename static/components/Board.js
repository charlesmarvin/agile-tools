import Card from './Card.js'
import CopyLink from './CopyLink.js'

const cardsByType = {
  standard: [0, '1/2', 1, 2, 3, 5, 8, 13, 20, 40, 100, '?'],
  fibonacci: [0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, '?'],
  tshirt: ['XS', 'S', 'M', 'L', 'XL', 'XXL', '?']
}

export default {
  name: 'Board',
  created: function () {
    const vm = this
    vm.boardId = this.$route.params.id
    vm.loadBoard(vm.boardId)
    vm.ws = new WebSocket(`ws://${window.location.host}/ws`)
    vm.ws.onopen = function (event) {
      console.log('connected! ', event)
      vm.ws.send(JSON.stringify({ type: 'init', request: { boardId: vm.boardId } }))
    }
    vm.ws.onclose = function (event) {
      console.log('disconnected ', event)
      vm.ws = null
    }
    vm.ws.onmessage = function (event) {
      console.log('msg: ', event.data)
      const msg = JSON.parse(event.data)
      switch (msg.type) {
        case 'init':
          vm.onInit(msg.response.sessionId)
          break
        case 'member_joined':
        case 'member_left':
          vm.members = msg.response.members
          break
        case 'vote_reveal':
          vm.votes = Object.values(msg.response.votes)
            .reduce((acc, it) => {
              acc[it] = acc[it] + 1 || 1
              return acc
            }, {})
          break
        default:
          break
      }
    }
  },
  components: {
    Card,
    CopyLink
  },
  computed: {
    link: function () {
      return window.location.href
    }
  },
  data: function () {
    return {
      boardId: null,
      activeBoard: null,
      selection: null,
      sessionId: null,
      members: [],
      votes: {},
      cardsByType
    }
  },
  methods: {
    onReset: function () {
      const vm = this
      vm.selection = null
    },
    onInit: function (sessionId) {
      const vm = this
      vm.sessionId = sessionId
    },
    loadBoard: async function (boardId) {
      console.log('Attempting to load board: ', boardId)
      if (!boardId) return
      const vm = this
      return fetch('/api/v1/boards/' + boardId, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res => res.ok ? res.json() : Promise.resolve(undefined))
        .then(response => { vm.activeBoard = response })
        .catch(error => { this.$root.$emit('alert-error', 'Oh snap. We had some issues loading that board.') })
    },
    vote: function () {
      const vm = this
      if (vm.selection) {
        vm.ws.send(JSON.stringify({
          type: 'vote',
          request: {
            boardId: vm.boardId,
            vote: `${vm.selection}`,
            sessionId: vm.sessionId
          }
        })
        )
      }
    }
  },
  template: `
    <div v-if="activeBoard">
      <div class="Board__header">
        <h2>{{ activeBoard.name }}</h2>
        <CopyLink v-bind:target="link"></CopyLink>
      </div>
      <div v-if=" members.length == 1">
        You are the only one on this board. Copy and share the link! 
      </div>
      <div v-if=" members.length > 1">
        {{ members.length }} Participants 
      </div>
      <div v-else>&nbsp;</div>
      <div class="Board__cards">
        <Card 
          v-for="card in cardsByType[activeBoard.type]" 
          v-bind:key="card" 
          v-bind:label="card"
          v-bind:selected="card  == selection"
          v-bind:votes="votes[card]"
          @click.native="selection = (card != selection) ? card : null"
        ></Card>
      </div>
      
      <div class="Board__actions">
        <div class="col">
          <button @click.prevent="vote" class="btn" type="submit">Vote</button>
        </div>
      </div>
    </div>
    <div v-else>
      <span>Board Not Found. <a href="/">Create or join a board.</a></span>
    </div>
  `
}
