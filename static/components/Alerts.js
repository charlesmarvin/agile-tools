export default {
  name: 'Alerts',
  created: function () {
    const vm = this
    this.$root.$on('alert-success', (msg) => vm.showAlert('success', msg))
    this.$root.$on('alert-info', (msg) => vm.showAlert('info', msg))
    this.$root.$on('alert-warn', (msg) => vm.showAlert('warn', msg))
    this.$root.$on('alert-error', (msg) => vm.showAlert('error', msg))
  },
  data: function () {
    return {
      alertClass: '',
      message: '',
      show: false
    }
  },
  methods: {
    close: function () {
      this.show = false
    },
    showAlert: function (type, message) {
      const vm = this
      this.alertClass = `alert-${type}`
      this.message = message
      this.show = true
      setTimeout(() => { vm.show = false }, 5000)
    }
  },
  template: `
    <div 
      class="alert alert-dismissible" 
      role="alert"
      v-bind:class="[alertClass]"
      v-if="show"
    >
      {{ message }}
      <button 
        aria-label="Close" 
        type="button" 
        class="close" 
        data-dismiss="alert" 
        v-on:click="close"
        >
        <span aria-hidden="true">&times;</span>
      </button>
    </div>
  `
}
