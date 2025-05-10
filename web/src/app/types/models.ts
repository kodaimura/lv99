export type Account = {
  id: number,
  name: string,
  role: number,
  created_at: string,
  updated_at: string,
}

export type Question = {
  id: number;
  title: string;
  content: string;
  answer: string;
  level: number;
  updated_at: string;
  deleted_at: string | null;
};

export type Answer = {
  answer_id: number;
  question_id: number;
  account_id: number;
  code_def: string;
  code_call: string;
  call_output: string;
  call_error: string;
  is_correct: boolean | null;
  correct_at: string | null;
  updated_at: string;
};

export type Comment = {
  comment_id: number;
  answer_id: number;
  account_id: number;
  comment_content: string;
  created_at: string;
  updated_at: string;
};