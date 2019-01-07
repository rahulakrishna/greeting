export class Poem {
  title: string;
  body: string[];
  author: string;

  constructor(title?: string, body?: string[], author?: string) {
    this.title = title || 'Solitary Reaper';
    this.body = body || [
      'Behold her, single in the field, ',
      'Yon solitary Highland Lass! ',
      'Reaping and singing by herself; ',
      'Stop here, or gently pass! ',
      'Alone she cuts and binds the grain, ',
      'And sings a melancholy strain; ',
      'O listen! for the Vale profound ',
      'Is overflowing with the sound. '
    ];
    this.author = author || 'William Wordsworth';
  }
}
