(function () {

    angular.module('selfieRoulette').factory('countdownTimerService', countdownTimerService);

    function countdownTimerService($timeout) {

        var self = this;
        self.start = start;
        self.stop = stop;
        self.reset = reset;
        self.getCount = getCount;
        self.setCount = setCount;

        var stopped;
        var counter = {count: 10};


        return {
            start: self.start,
            stop: self.stop,
            reset: self.reset,
            getCount: self.getCount,
            setCount: self.setCount
        };


        function start() {
            stopped = $timeout(function () {
                if (counter.count-- > 0) {
                    self.start();
                }
            }, 1000);
        }


        function stop() {
            $timeout.cancel(stopped);
            reset();
        }

        function reset() {
            counter.count = 10;
        }

        function getCount() {
            return counter;
        }

        function setCount(countValue) {
            counter.count = countValue;
        }

    }

})();
