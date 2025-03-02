# Analysis Service

## Описание сервиса

Analysis Service - это аналитический сервис, предназначенный для прогнозирования вероятности
дефолта и восстановления фондового рынка с использованием фрактальной модели.
Сервис применяет как линейные, так и нелинейные подходы к анализу различных факторов,
влияющих на состояние рынка.

### Ключевые особенности

1. Фрактальная модель: Анализирует сложные зависимости между различными параметрами рынка,
используя фрактальные характеристики.
2. Нелинейная модель: Включает полиномиальные термины для каждого фактора, 
что позволяет учесть более сложные зависимости.
3. Многомерный анализ: Учитывает технические индикаторы (RSI, SMA), 
новостные настроения, политическую ситуацию и другие ключевые факторы.
4. Гибкость: Возможность адаптации модели под конкретные условия рынка 
через настройку коэффициентов и добавление новых переменных.

### Математическая модель

#### Линейная модель

Линейная модель определяется следующим уравнением:

[//]: # ($$)

[//]: # (P_{default} = \alpha_1 H + \alpha_2 RSI + \alpha_3 SMA + \alpha_4 St + \alpha_5 P)

[//]: # ($$)
![](https://latex.codecogs.com/svg.image?{\color{Red}P_{default}=\alpha_1&space;H&plus;\alpha_2&space;RSI&plus;\alpha_3&space;SMA&plus;\alpha_4&space;St&plus;\alpha_5&space;P})

Где:
* H — фрактальная характеристика,
* RSI — относительная сила индикатора,
* SMA — скользящая средняя,
* St — новостные настроения,
* ![](https://latex.codecogs.com/svg.image?{\color{Red}P_{political}}) - политические настроения.
* ![](https://latex.codecogs.com/svg.image?{\color{Red}\alpha_1,\alpha_2,\alpha_3,\alpha_4,\alpha_5}) — коэффициенты, 
оцениваемые методом наименьших квадратов или другими методами оптимизации. 

#### Нелинейная модель

Нелинейная модель включает дополнительные преобразования для учета нелинейных зависимостей:

![](https://latex.codecogs.com/svg.image?{\color{Red}P_{default}=\alpha_1&space;H^2&plus;\alpha_2\cdot\tanh(RSI)&plus;\alpha_3\cdot\log(SMA)&plus;\alpha_4\cdot&space;e^{St}&plus;\alpha_5\cdot\frac{1}{1&plus;e^{-P_{political}}}})

#### Объяснение терминов:

* ![](https://latex.codecogs.com/svg.image?{\color{Red}H^2}): Фрактальная характеристика,
возведенная в квадрат,
для усиления её влияния на вероятность дефолта.
* ![](https://latex.codecogs.com/svg.image?{\color{Red}tanh(RSI)}): Гиперболический тангенс для изменения поведения RSI 
в диапазоне от -1 до 1, что делает модель более гибкой при больших значениях RSI.
* ![](https://latex.codecogs.com/svg.image?{\color{Red}log(SMA)tanh(RSI)}): Логарифм скользящей средней для замедления 
её влияния при очень больших значениях.
* ![](https://latex.codecogs.com/svg.image?{\color{Red}_e{St}}): Экспоненциальная зависимость от новостных настроений,
учитывающая сильное влияние экстремальных значений настроений.
* ![](https://latex.codecogs.com/svg.image?{\color{Red}\frac{1}{1&plus;e^{-P_{political}}}): Сигмоидальная функция для 
ограничения влияния политической нестабильности в пределах от 0 до 1.
* ![](https://latex.codecogs.com/svg.image?{\color{Red}\alpha_1,\alpha_2,\alpha_3,\alpha_4,\alpha_5}) также оцениваются
на основе исторических данных.

