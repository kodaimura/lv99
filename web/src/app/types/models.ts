export type Question = {
  question_id: number;
  question_title: string;
  question_content: string;
  question_answer: string;
  question_level: number;
  updated_at: string;
  deleted_at: string | null;
};