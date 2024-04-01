export type IQuiz = {
  id: number,
  name: string,
  preview: string,
}

export type IQuizQuestion = {
  id: number,
  name: string,
  preview: string,
  priority: number,
}

export type IQuizGame = {
  id: number,
  name: string,
  preview: string,
  options: IQuizQuestion[],
}

export type IQuizResults = {
  quizId: number,
  percentageWins: number,
}