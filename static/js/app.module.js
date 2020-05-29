import { CreateBoardForm } from './create-board-form.module.js'
export class App {
  static start () {
    console.log('Starting application...')
    const createBoardForm = new CreateBoardForm()
    createBoardForm.init('create-board-form')
  }
}
