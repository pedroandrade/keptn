import * as moment from "moment";

export default class DateUtil {
  static getDurationFormatted(start, end) {
    let diff = moment(start).diff(moment(end));
    let duration = moment.duration(diff);

    let result = moment.utc(diff).format("s")+' seconds';

    if(duration.asMinutes() > 0)
      result = moment.utc(diff).format("mm")+' minutes '+result;

    if(duration.asHours() > 0)
      result = Math.floor(duration.asHours())+' hours '+result;

    return result;
  }

  static getCalendarFormats(showSeconds?: boolean) {
    if(showSeconds) {
      return {
        lastDay : '[Yesterday at] HH:mm:ss',
        sameDay : '[Today at] HH:mm:ss',
        nextDay : '[Tomorrow at] HH:mm:ss',
        lastWeek : '[last] dddd [at] HH:mm:ss',
        nextWeek : 'dddd [at] HH:mm:ss',
        sameElse : 'YYYY-MM-DD HH:mm:ss'
      };
    }
    return {
      lastDay : '[Yesterday at] HH:mm',
      sameDay : '[Today at] HH:mm',
      nextDay : '[Tomorrow at] HH:mm',
      lastWeek : '[last] dddd [at] HH:mm',
      nextWeek : 'dddd [at] HH:mm',
      sameElse : 'YYYY-MM-DD HH:mm'
    };
  }
}
