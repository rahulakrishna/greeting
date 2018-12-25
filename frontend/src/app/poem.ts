export class Poem {
  title: string;
  body: string[];
  author: string;

  constructor(title?: string, body?: string[], author?: string) {
    this.title = title || "സൂര്യകാന്തി";
    this.body = body || [
      "മന്ദമന്ദമെന്‍ താഴും മുഗ്ദമാം മുഖം പൊക്കി-",
      "സ്സുന്ദരദിവാകരന്‍ ചോദിച്ചൂ മധുരമായ്‌:",
      "“ആരു നീയനുജത്തീ? നിര്‍ന്നിമേഷയായെന്തെന്‍",
      "തേരുപോകവെ നേരെ നോക്കിനില്‍ക്കുന്നൂ ദൂരേ?",
      "സൗമ്യമായ്‌ പിന്നെപ്പിന്നെ വിടരും കണ്ണാല്‍ സ്നേഹ-",
      "രമ്യമായ്‌ വീക്ഷിയ്ക്കുന്നൂ തിരിഞ്ഞു തിരിഞ്ഞെന്നെ;",
      "വല്ലതും പറയുവാനാഗ്രഹിയ്ക്കുന്നുണ്ടാവാ-",
      "മില്ലയോ? തെറ്റാണൂഹമെങ്കിൽ, ഞാന്‍ ചോദിച്ചീല.”",
      "ഒന്നുമുത്തരം തോന്നീലെങ്ങനെ തോന്നും? സര്‍വ്വ-",
      "സന്നുതന്‍ സവിതാവെങ്ങു നിര്‍ഗന്ധം പുഷ്പം!",
      "അര്യമാവിനെ സ്നേഹിക്കുന്ന ധിക്കാരത്തിന്നു",
      "സൂര്യകാന്തിയെന്നെന്നെ പ്പുച്ഛിച്ചതാണീ ലോകം!"
    ];
    this.author = author || "ജി. ശങ്കരക്കുറുപ്പ്";
  }
}