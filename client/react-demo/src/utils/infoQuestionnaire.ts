import { Questionnaire } from '../types/questionnaire';

const InfoQuestionnaire: Questionnaire = {
  id: 'user_information_questionnaire',
  created_at: 'dne',
  name: '使用者基本資料',
  questions: [
    {
      id: 'birth_year',
      body: '出生年(民國)',
      type: 'short_answer',
    },
    {
      id: 'height',
      body: '身高(公分)',
      type: 'short_answer',
    },
    {
      id: 'weight',
      body: '體重(公斤)',
      type: 'short_answer',
    },
    {
      id: 'gender',
      body: '性別',
      type: 'single_choice',
      choices: [
        { id: 'male', body: '男' },
        { id: 'female', body: '女' },
        { id: 'nonbinary', body: '中性' },
      ],
    },
    {
      id: 'education_level',
      body: '教育程度',
      type: 'single_choice',
      choices: [
        { id: 'elementry_school_or_below', body: '小學或以下' },
        { id: 'middle_school', body: '國中' },
        { id: 'high_school', body: '高中' },
        { id: 'bachelor', body: '大學' },
        { id: 'master', body: '碩士' },
        { id: 'doctorate', body: '博士' },
      ],
    },
    {
      id: 'occupation',
      body: '職業',
      type: 'single_choice',
      choices: [
        { id: 'student', body: '學生' },
        { id: 'government_employee', body: '軍公教' },
        { id: 'service_industry', body: '服務業' },
        { id: 'industry_and_commerce', body: '工商業' },
        { id: 'freelancer', body: '自由業' },
        { id: 'domestic', body: '家管' },
        { id: 'retired', body: '退休人員' },
        { id: 'others', body: '其他' },
      ],
    },
    {
      id: 'marriage',
      body: '婚姻狀況',
      type: 'single_choice',
      choices: [
        { id: 'single', body: '未婚' },
        { id: 'married', body: '已婚' },
        { id: 'divorced', body: '離婚' },
        { id: 'widowed', body: '喪偶' },
      ],
    },
    {
      id: 'medical_history',
      body: '疾病史',
      type: 'single_choice',
      choices: [
        { id: 'high_blood_pressure', body: '高血壓' },
        { id: 'hyperlipidemia', body: '高血脂(高膽固醇)' },
        { id: 'diabetes', body: '糖尿病' },
        { id: 'heart_disease', body: '心臟病' },
        { id: 'stroke', body: '腦中風' },
        { id: 'mental_illness', body: '精神疾病' },
        { id: 'dementia', body: '失智症' },
        { id: 'none_of_the_above', body: '沒有上面提到的疾病' },
      ],
    },
    {
      id: 'medication_status',
      body: '服用藥物狀況',
      type: 'single_choice',
      choices: [
        { id: 'cardiovascular_drugs', body: '三高類型或心血管疾病藥物' },
        { id: 'psychiatric_drugs', body: '精神疾病藥物' },
        { id: 'other_drugs', body: '其他疾病藥物' },
        { id: 'no_drugs', body: '無使用藥物' },
      ],
    },
    {
      id: 'demented_among_direct_relatives',
      body: '直系血親(雙親、外公外婆與阿公阿嬤)中是否有人罹患失智?',
      type: 'single_choice',
      choices: [
        { id: 'false', body: '無' },
        { id: 'true', body: '有' },
      ],
    },
    {
      id: 'head_injury_experience',
      body: '是否有過頭部外傷的經驗',
      type: 'single_choice',
      choices: [
        { id: 'false', body: '沒有' },
        { id: 'true', body: '有' },
      ],
    },
    {
      id: 'ear_condition',
      body: '聽力狀況',
      type: 'single_choice',
      choices: [
        { id: 'normal', body: '正常' },
        { id: 'slightly_affecting_conversation', body: '有點影響溝通' },
        { id: 'need_hearing_aid', body: '需使用助聽器才聽得見' },
      ],
    },
    {
      id: 'eyesight_condition',
      body: '視力狀況',
      type: 'single_choice',
      choices: [
        { id: 'normal', body: '正常' },
        { id: 'slightly_affecting_reading', body: '有點影響閱讀' },
        { id: 'need_glasses', body: '需配戴眼鏡才看得見' },
      ],
    },
    {
      id: 'smoking_habit',
      body: '抽菸的習慣',
      type: 'single_choice',
      choices: [
        { id: 'none', body: '沒有' },
        { id: 'sometimes', body: '偶而抽菸' },
        { id: 'everyday', body: '每天抽菸' },
      ],
    },
  ],
};

export default InfoQuestionnaire;
