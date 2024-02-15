import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const getDurationFromNow = (date: Date) => {
  // minute, hour, day, week, month, year
  // return: 2 minutes ago, 1 hour ago, 1 day ago, 1 week ago, 1 month ago, 1 year ago
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  const seconds = Math.floor(diff / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);
  const weeks = Math.floor(days / 7);
  const months = Math.floor(days / 30);
  const years = Math.floor(days / 365);

  if (seconds < 60) {
    if (seconds == 1) {
      return `${seconds} second ago`;
    } else {
      return `${seconds} seconds ago`;
    }
  } else if (minutes < 60) {
    if (minutes == 1) {
      return `${minutes} minute ago`;
    } else {
      return `${minutes} minutes ago`;
    }
  } else if (hours < 24) {
    if (hours == 1) {
      return `${hours} hour ago`;
    } else {
      return `${hours} hours ago`;
    }
  } else if (days < 7) {
    if (days == 1) {
      return `${days} day ago`;
    } else {
      return `${days} days ago`;
    }
  } else if (weeks < 4) {
    if (weeks == 1) {
      return `${weeks} week ago`;
    } else {
      return `${weeks} weeks ago`;
    }
  } else if (months < 12) {
    if (months == 1) {
      return `${months} month ago`;
    } else {
      return `${months} months ago`;
    }
  } else {
    if (years == 1) {
      return `${years} year ago`;
    } else {
      return `${years} years ago`;
    }
  }
};
