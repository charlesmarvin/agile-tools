export default {
  name: 'CreateBoardForm',
  data: function () {
    return {
      boardName: '',
      boardType: 'standard',
      passcode: '',
      passcodeProtect: false,
      passcodeLength: 4
    }
  },
  methods: {
    createBoard: function (event) {
      var vm = this
      var payload = {
        name: this.boardName,
        type: this.boardType
      }
      if (this.passcodeProtect) {
        payload.passcode = this.passcode
      }
      console.log('creating board ', payload)
      if (!payload.name || !payload.type) {
        vm.$root.$emit('alert-error', 'Name and type are required.')
        return
      }
      fetch('/api/v1/boards', {
        method: 'POST',
        body: JSON.stringify(payload),
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res => res.json())
        .then(res => {
          window.location.hash = res.id
          vm.show = false
          vm.$root.$emit('alert-success', 'Cool. Your board was created.')
          this.$router.push({ name: 'board', params: { id: res.id } })
        })
        .catch(error => { vm.$root.$emit('alert-error', 'Well that\'s embarrassing. Looks like we messed up creating your board.') })
    },
    toggleSecure: function () {
      this.passcodeProtect = !this.passcodeProtect
    }
  },
  template: `
    <form id="create-board-form">
      <div class="flex-grid toggle">
        <label v-if="passcodeProtect" v-on:click="toggleSecure">
          <span class="icon svg-baseline text-large">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
              <path fill-rule="evenodd" d="M7,11 L7,6 C7,3.23857625 9.23857625,1 12,1 C14.7614237,1 17,3.23857625 17,6 L17,8 L15,8 L15,6 C15,4.34314575 13.6568542,3 12,3 C10.3431458,3 9,4.34314575 9,6 L9,11 L18,11 C19.1045695,11 20,11.8954305 20,13 L20,21 C20,22.1045695 19.1045695,23 18,23 L6,23 C4.8954305,23 4,22.1045695 4,21 L4,13 C4,11.8954305 4.8954305,11 6,11 L7,11 Z M6,13 L6,21 L18,21 L18,13 L6,13 Z M12,18 C11.4477153,18 11,17.5522847 11,17 C11,16.4477153 11.4477153,16 12,16 C12.5522847,16 13,16.4477153 13,17 C13,17.5522847 12.5522847,18 12,18 Z"/>
            </svg>
          </span>
        </label>
        <label v-else v-on:click="toggleSecure">
          <span class="icon svg-baseline text-large">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
              <path fill-rule="evenodd" d="M7,10 L7,7 C7,4.23857625 9.23857625,2 12,2 C14.7614237,2 17,4.23857625 17,7 L17,10 L18,10 C19.0683513,10 20,10.7763739 20,11.8333333 L20,20.1666667 C20,21.2236261 19.0683513,22 18,22 L6,22 C4.93164867,22 4,21.2236261 4,20.1666667 L4,11.8333333 C4,10.7763739 4.93164867,10 6,10 L7,10 Z M9,10 L15,10 L15,7 C15,5.34314575 13.6568542,4 12,4 C10.3431458,4 9,5.34314575 9,7 L9,10 Z M6,12 L6,20 L18,20 L18,12 L6,12 Z M12,17 C11.4477153,17 11,16.5522847 11,16 C11,15.4477153 11.4477153,15 12,15 C12.5522847,15 13,15.4477153 13,16 C13,16.5522847 12.5522847,17 12,17 Z"/>
            </svg>
          </span>
        </label>
      </div>
      <div class="flex-grid-sp thirds sp-bottom">
        <label class="col board-type-select">
          <input type="radio" name="boardType" v-model="boardType" value="standard">
          <span>Standard</span>
        </label>
        <label class="col board-type-select">
          <input type="radio" name="boardType" v-model="boardType" value="fibonacci">
          <span>Fibonacci</span>
        </label>
        <label class="col board-type-select">
          <input type="radio" name="boardType" v-model="boardType" value="tshirt">
          <span>T-Shirt</span>
        </label>
      </div>
      
      <div class="flex-grid sp-bottom">
        <div class="col">
          <input v-model="boardName" type="text" name="boardName" placeholder="Sprint Planning 123" id="board-name-input">
        </div>
      </div>
      
      <div class="flex-grid sp-bottom" v-bind:class="{ hide: !passcodeProtect }">
        <div class="col">
          <input v-model="passcode" type="text" name="passcode" placeholder="4-digit Passcode" id="board-passcode-input" :maxlength="passcodeLength">
        </div>
      </div>

      <div class="flex-grid-sp">
        <div class="col">
          <button @click.prevent="createBoard" class="btn" type="submit">Create Board</button>
        </div>
      </div>
    </form>
  `
}
