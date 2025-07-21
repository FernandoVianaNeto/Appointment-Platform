import { useState } from 'react';
import dayjs, { Dayjs } from 'dayjs';
import { CircleButton, DateContainer, TodayText, Wrapper } from './styles';

interface IDateSelectorProps {
  onClick: (day: Dayjs) => void;
}

function DateSelector({ onClick }: IDateSelectorProps) {
  const [selectedDate, setSelectedDate] = useState(dayjs());

  function goToPrevDay() {
    const newDate = selectedDate.subtract(1, 'day');
    setSelectedDate(newDate);
    onClick?.(newDate);
  }
  
  function goToNextDay() {
    const newDate = selectedDate.add(1, 'day');
    setSelectedDate(newDate);
    onClick?.(newDate);
  }
  
  function resetToToday() {
    const today = dayjs();
    setSelectedDate(today);
    onClick?.(today);
  }

  return (
    <Wrapper>
      <CircleButton onClick={goToPrevDay}>‹</CircleButton>
      <DateContainer>
        <strong>{selectedDate.format('MMM D, YYYY')}</strong>
        <TodayText onClick={resetToToday}>
          &nbsp;Today ▾
        </TodayText>
      </DateContainer>
      <CircleButton onClick={goToNextDay}>›</CircleButton>
    </Wrapper>
  );
}


export default DateSelector;