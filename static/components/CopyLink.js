export default {
  name: 'CopyLink',
  props: {
    target: String
  },
  data: function () {
    return {
      copied: false
    }
  },
  methods: {
    copyToClipboard: function () {
      const vm = this
      const link = document.querySelector('#board-link')
      link.setAttribute('type', 'text')
      link.select()

      try {
        if (document.execCommand('copy')) {
          vm.copied = true
          setTimeout(() => {
            vm.copied = false
          }, 800)
        }
      } catch (err) {
        alert('Oops, unable to copy')
      }
      link.setAttribute('type', 'hidden')
      window.getSelection().removeAllRanges()
    }
  },
  template: `
    <div>
      <span v-if="copied">Copied!</span>
      <a v-else v-bind:title="target" @click.prevent="copyToClipboard">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M14,9 L14,7 L18,7 C20.7614237,7 23,9.23857625 23,12 C23,14.7614237 20.7614237,17 18,17 L14,17 L14,15 L18,15 C19.6568542,15 21,13.6568542 21,12 C21,10.3431458 19.6568542,9 18,9 L14,9 Z M10,15 L10,17 L6,17 C3.23857625,17 1,14.7614237 1,12 C1,9.23857625 3.23857625,7 6,7 L10,7 L10,9 L6,9 C4.34314575,9 3,10.3431458 3,12 C3,13.6568542 4.34314575,15 6,15 L10,15 Z M7,13 L7,11 L17,11 L17,13 L7,13 Z"/></svg>
      </a>
      <input type="hidden" id="board-link" :value="target">
    </div>
  `
}
