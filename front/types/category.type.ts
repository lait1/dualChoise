import type {IQuiz} from "#build/types/quiz.type";

export type ICategoryInfo = {
  id: number,
  name: string,
  preview: string,
  created: string,
  countQuizzes: number,
}
export type ICategory = {
  categoryId: number,
  categoryName: string,
  quizzes: IQuiz[],
}