export interface Choice {
  id: string;
  body: string;
}

export interface Question {
  id: string;
  type: 'short_answer' | 'single_choice';
  body: string;
  choices?: Choice[];
}

export interface Answer {
  question_id: string;
  type: 'short_answer' | 'single_choice';
  body: string;
}

export interface Questionnaire {
  id: string;
  name: string;
  created_at: string;
  questions: Question[];
}
