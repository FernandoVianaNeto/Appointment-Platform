import { useState } from 'react';
import dayjs from 'dayjs';
import { CircleButton, DateContainer, TodayText, Wrapper } from './styles';

function DateSelector() {
  const [selectedDate, setSelectedDate] = useState(dayjs());

  function goToPrevDay() {
    setSelectedDate((prev) => prev.subtract(1, 'day'));
  }

  function goToNextDay() {
    setSelectedDate((prev) => prev.add(1, 'day'));
  }

  function resetToToday() {
    setSelectedDate(dayjs());
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