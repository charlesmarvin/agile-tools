export default {
  name: 'Home',
  template: `
    <div class="Home">
      <nav>
        <router-link class="Home__link" :to="{ name: 'board'}">Join Existing Board</router-link>
        <router-link class="Home__link" :to="{ name: 'create-board'}">Create New Board</router-link>
      </nav>
    </div>
  `
}
