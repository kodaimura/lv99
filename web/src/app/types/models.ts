export type Question = {
  question_id: number;
  question_title: string;
  question_content: string;
  question_answer: string;
  question_level: number;
  updated_at: string;
  deleted_at: string | null;
};

export type Answer = {
  answer_id: number;
  question_id: number;
  code_def: string;
  code_call: string;
  is_correct: boolean | null;
  correct_at: string | null;
  updated_at: string;
};